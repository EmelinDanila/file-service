package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/EmelinDanila/file-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFileServiceClient(conn)

	// Upload file
	uploadFile(c, "test.jpg")

	// List files
	listFiles(c)

	// Download file
	downloadFile(c, "test.jpg")
}

func uploadFile(c pb.FileServiceClient, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stream, err := c.UploadFile(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := stream.Send(&pb.FileChunk{
			Filename: filename,
			Content:  buffer[:n],
		}); err != nil {
			log.Fatal(err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Upload response: %s\n", resp.Message)
}

func listFiles(c pb.FileServiceClient) {
	resp, err := c.ListFiles(context.Background(), &pb.ListRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name | Created At | Updated At")
	for _, file := range resp.Files {
		fmt.Printf("%s | %s | %s\n",
			file.Filename,
			time.Unix(file.CreatedAt, 0).Format(time.RFC1123),
			time.Unix(file.UpdatedAt, 0).Format(time.RFC1123),
		)
	}
}

func downloadFile(c pb.FileServiceClient, filename string) {
	stream, err := c.DownloadFile(context.Background(), &pb.FileRequest{Filename: filename})
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("downloaded_" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		output.Write(chunk.Content)
	}
	fmt.Printf("File %s downloaded\n", filename)
}
