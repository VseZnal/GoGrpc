package main

import (
	protodeletemem "Grpc/services/deletemem/protodeletemem/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	protodeletemem.UnimplementedDeletememServiceServer
}

func main() {
	//e := protodeletemem.Init()
	//if e != nil {
	//	log.Fatal(e)
	//}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protodeletemem.RegisterDeletememServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
