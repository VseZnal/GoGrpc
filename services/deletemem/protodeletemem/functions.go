package protodeletemem

import (
	postgresql "Grpc/services/deletemem/database"
	"Grpc/services/deletemem/protodeletemem/proto"
	pb "Grpc/services/deletemem/protodeletemem/proto"
	"context"
	"fmt"
)

type Server struct {
	pb.UnimplementedDeletememServiceServer
}

var db postgresql.Database

func Init() error {
	var err error

	db, err = postgresql.NewDatabase()
	return err
}

func (s Server) Delete(ctx context.Context, request *protodeletemem.DeleteMemRequest) (*protodeletemem.DeleteMemResponse, error) {
	err := db.Deletedata_db(request.Slug)

	if err != nil {
		fmt.Printf("enteded delete msg info")
	}
	return &protodeletemem.DeleteMemResponse{Status: 1}, nil
}
