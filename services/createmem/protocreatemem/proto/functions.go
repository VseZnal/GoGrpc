package protocreatemem

import (
	postgresql "Grpc/services/createmem/database"
	"Grpc/services/createmem/models"
	"context"
)

type Server struct {
	Port string
}

var db postgresql.Database

func Init() error {
	var err error

	db, err = postgresql.NewDatabase()
	return err
}

func (s Server) Create(ctx context.Context, request *CreateMemRequest) (*CreateMemResponse, error) {
	msg := models.CreateReq{
		ParentId: request.ParentId,
		UserId:   request.UserId,
		PostId:   request.PostId,
		Content:  request.Content}

	_ = db.Postdata_db(msg)

	return &CreateMemResponse{}, nil
}
