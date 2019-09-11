package main

import (
	"fmt"
	"github.com/mp3tags/request-repository-proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"request-repository/service"
)

func main() {
	srv := service.New()
	srv.Db = service.ConnectToDb()
	startServer(srv)
}

func startServer(service *service.Service) {
	// create a TCP listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Println("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	request_repository.RegisterRequestRepositoryServiceServer(grpcServer, service)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve: %s", err)
	}
}
