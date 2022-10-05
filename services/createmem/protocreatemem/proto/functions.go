package protocreatemem

import (
	"context"
)

type Server struct {
	Port string
}

func (s Server) Create(ctx context.Context, request *CreateMemRequest) (*CreateMemResponse, error) {
	//TODO implement me
	panic("implement me")
}
