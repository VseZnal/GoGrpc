package protocreatemem

import (
	postgresql "Grpc/services/createmem/database"
	"Grpc/services/createmem/models"
	pb "Grpc/services/createmem/protocreatemem/proto"
	protocreatemem "Grpc/services/createmem/protocreatemem/proto"
	"context"
	"fmt"
)

type Server struct {
	pb.UnimplementedCreatememServiceServer
}

var db postgresql.Database

func Init() error {
	var err error

	db, err = postgresql.NewDatabase()
	return err
}

func (s Server) Create(ctx context.Context, request *protocreatemem.CreateMemRequest) (*protocreatemem.CreateMemResponse, error) {
	msg := models.CreateReq{
		ParentId: request.ParentId,
		UserId:   request.UserId,
		PostId:   request.PostId,
		Content:  request.Content}

	err := db.Postdata_db(msg)

	if err != nil {
		fmt.Printf("enteted get msg info")
	}
	out := &protocreatemem.Memcreate{
		ParentId: msg.ParentId,
		UserId:   msg.UserId,
		PostId:   msg.PostId,
		Content:  msg.Content,
	}

	return &protocreatemem.CreateMemResponse{Memcreate: out}, nil
}
