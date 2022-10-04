package main

import (
	"Grpc/services/getmem/app"
	protogetmem "Grpc/services/getmem/protogetmem/proto"
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
		grpc.UnaryInterceptor(protogetmem.AccessLogInterceptor),
	)
	s := &protogetmem.Server{}

	// attach the user service to the server
	protogetmem.RegisterGetmemServiceServer(grpcServer, s)

	log.Printf("%s service started on  %s:%s (localhos:%s)", app.SERVICE_NAME, Host, Port, LocalPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
