package protodeletemem

import (
	"context"
)

type Server struct {
	Port string
}

func (s Server) Delete(ctx context.Context, request *DeleteMemRequest) (*DeleteMemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedDeletememServiceServer() {
	//TODO implement me
	panic("implement me")
}
