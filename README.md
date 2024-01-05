#to create go file fromm proto
protoc -I. --go_out=plugins=grpc:. trisa.proto

#builds 
docker build -t trisa-server .
docker run -p 50051:50051 trisa-server
