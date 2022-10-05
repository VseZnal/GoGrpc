package main

import (
	protocreatemem "Grpc/services/createmem/protocreatemem/proto"
	"google.golang.org/grpc"

	"log"
	"net"
)

type server struct {
	protocreatemem.UnimplementedCreatememServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocreatemem.RegisterCreatememServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
