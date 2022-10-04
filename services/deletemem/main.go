package main

import (
	"Grpc/services/deletemem/app"
	protodeletemem "Grpc/services/deletemem/protodeletemem/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	Host := os.Getenv("HOST")
	Port := os.Getenv("PORT")
	LocalPort := os.Getenv("PORT_LOCAL")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(protodeletemem.AccessLogInterceptor),
	)
	s := &protodeletemem.Server{}

	// attach the user service to the server
	protodeletemem.RegisterDeletememServiceServer(grpcServer, s)

	log.Printf("%s service started on  %s:%s (localhos:%s)", app.SERVICE_NAME, Host, Port, LocalPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
