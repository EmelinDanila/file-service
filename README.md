# File Service

File Service — это gRPC-сервис, написанный на Go, который позволяет загружать, скачивать и просматривать список бинарных файлов (e.g. изображения). Сервис ограничивает количество одновременных подключений и сохраняет файлы на диск.

## Функциональность
- **Загрузка файлов:** прием бинарных файлов от клиента и сохранение в `storage`.
- **Скачивание файлов:** отдача файла по запросу.
- **Список файлов:** возвращает имя, дату создания и обновления для каждого файла.
- **Ограничения на количество запросов:**
  - 10 на загрузку/скачивание.
  - 100 на список файлов.

## Требования
- Go 1.21+
- protoc 3.21+
- `protoc-gen-go`, `protoc-gen-go-grpc`

## Установка
1. Клонируйте репозиторий:
```bash
git clone https://github.com/EmelinDanila/file-service.git
cd file-service
```
2. Установка зависимостей:
```bash
go mod tidy
```
3. Установите `protoc` и плагины:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
4. Сгенерируйте `.go` из `.proto`:
```bash
cd file-service
protoc --go_out=. --go-grpc_out=. proto/file_service.proto
```

## Структура
```
file-service/
├── go.mod
├── proto/
│   ├── file_service.proto
│   ├── file_service.pb.go
│   └── file_service_grpc.pb.go
├── server/
│   └── main.go
├── client/
│   └── main.go

```

## Запуск
### Сервер
```bash
cd server
go run main.go
```
### Клиент
```bash
cd client
go run main.go
```

## Пример клиента
```
Upload response: File uploaded successfully
Name | Created At | Updated At
test.jpg | Thu, 10 Apr 2025 18:57:01 MSK | Thu, 10 Apr 2025 18:57:01 MSK
File test.jpg downloaded
```

## Примечания
- Файлы хранятся в `./storage` относительно сервера.
