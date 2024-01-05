FROM golang:latest

WORKDIR /app

# Install protoc compiler
RUN apt-get update && apt-get install -y protobuf-compiler

# Download and install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Copy the proto files to the container
COPY trisa.proto .

# Generate Go code from proto files
RUN protoc -I. --go_out=plugins=grpc:. trisa.proto

# Copy the rest of your application files
COPY . .

# Build the Go application
RUN go mod download
RUN go build -o main .

EXPOSE 50051

CMD ["./main"]
