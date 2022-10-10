package main

import (
	"Grpc/services/getmem/protogetmem"
	pb "Grpc/services/getmem/protogetmem/proto"

	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	e := protogetmem.Init()
	if e != nil {
		log.Fatal(e)
	}

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetmemServiceServer(s, &protogetmem.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

/*

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

*/
