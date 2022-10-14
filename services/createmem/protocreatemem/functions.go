package protocreatemem

import (
	postgresql "Grpc/services/createmem/database"
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
	out := &protocreatemem.Memcreate{
		Slug:     "77",
		ParentId: request.ParentId,
		UserId:   request.UserId,
		PostId:   request.PostId,
		Content:  request.Content,
		Status:   0,
	}

	err := db.Postdata_db(out)

	if err != nil {
		fmt.Printf("enteted get msg info")
	}

	return &protocreatemem.CreateMemResponse{Memcreate: out}, err
}
