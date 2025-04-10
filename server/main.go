package main

import (
	"context" // Добавьте этот импорт
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	pb "github.com/EmelinDanila/file-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
	uploadSem   chan struct{} // semaphore for upload/download
	listSem     chan struct{} // semaphore for list
	storagePath string
	mu          sync.Mutex
}

func NewServer() *server {
	return &server{
		uploadSem:   make(chan struct{}, 10),  // 10 concurrent upload/download
		listSem:     make(chan struct{}, 100), // 100 concurrent list
		storagePath: "./storage",
	}
}

func (s *server) UploadFile(stream pb.FileService_UploadFileServer) error {
	s.uploadSem <- struct{}{}        // acquire semaphore
	defer func() { <-s.uploadSem }() // release semaphore

	var filename string
	var file *os.File

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadResponse{Message: "File uploaded successfully"})
		}
		if err != nil {
			return err
		}

		if filename == "" {
			filename = chunk.Filename
			os.MkdirAll(s.storagePath, 0755)
			file, err = os.Create(filepath.Join(s.storagePath, filename))
			if err != nil {
				return err
			}
			defer file.Close()
		}

		_, err = file.Write(chunk.Content)
		if err != nil {
			return err
		}
	}
}

func (s *server) DownloadFile(req *pb.FileRequest, stream pb.FileService_DownloadFileServer) error {
	s.uploadSem <- struct{}{}        // acquire semaphore
	defer func() { <-s.uploadSem }() // release semaphore

	filePath := filepath.Join(s.storagePath, req.Filename)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.FileChunk{
			Filename: req.Filename,
			Content:  buffer[:n],
		}); err != nil {
			return err
		}
	}
}

func (s *server) ListFiles(ctx context.Context, _ *pb.ListRequest) (*pb.ListResponse, error) {
	s.listSem <- struct{}{}        // acquire semaphore
	defer func() { <-s.listSem }() // release semaphore

	files, err := os.ReadDir(s.storagePath)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListResponse{}
	for _, file := range files {
		info, err := os.Stat(filepath.Join(s.storagePath, file.Name()))
		if err != nil {
			continue
		}
		resp.Files = append(resp.Files, &pb.FileInfo{
			Filename:  file.Name(),
			CreatedAt: info.ModTime().Unix(),
			UpdatedAt: info.ModTime().Unix(),
		})
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, NewServer())

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
