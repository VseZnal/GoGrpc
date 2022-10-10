package main

import (
	"Grpc/services/createmem/protocreatemem"
	pb "Grpc/services/createmem/protocreatemem/proto"
	"google.golang.org/grpc"

	"log"
	"net"
)

func main() {

	e := protocreatemem.Init()
	if e != nil {
		log.Fatal(e)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCreatememServiceServer(s, &protocreatemem.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
