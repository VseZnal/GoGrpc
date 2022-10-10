package protogetmem

import (
	postgresql "Grpc/services/getmem/database"
	pb "Grpc/services/getmem/protogetmem/proto"
	protogetmem "Grpc/services/getmem/protogetmem/proto"
	"context"
	"fmt"
)

type Server struct {
	pb.UnimplementedGetmemServiceServer
}

var db postgresql.Database

func Init() error {
	var err error

	db, err = postgresql.NewDatabase()
	return err
}

func (s Server) Get(ctx context.Context, request *protogetmem.GetMemRequest) (*protogetmem.GetMemResponse, error) {
	msg, err := db.Getdata_db(request.Slug)

	out := &protogetmem.Memget{
		ParentId: msg.ParentId,
		UserId:   msg.UserId,
		PostId:   msg.PostId,
		Content:  msg.Content,
	}
	if err != nil {
		fmt.Printf("enteded get msg info")
	}
	return &protogetmem.GetMemResponse{Memget: out}, err

}
