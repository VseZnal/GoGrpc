package main

import (
	"Grpc/services/deletemem/protodeletemem"
	pb "Grpc/services/deletemem/protodeletemem/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	e := protodeletemem.Init()
	if e != nil {
		log.Fatal(e)
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDeletememServiceServer(s, &protodeletemem.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
