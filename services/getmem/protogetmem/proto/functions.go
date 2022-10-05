package protogetmem

import (
	"context"
)

type Server struct {
	Port string
}

func (s Server) Get(ctx context.Context, request *GetMemRequest) (*GetMemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedGetmemServiceServer() {
	//TODO implement me
	panic("implement me")
}
