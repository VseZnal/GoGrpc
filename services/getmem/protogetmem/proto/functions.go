package protogetmem

import (
	"Grpc/services/api-gw/app"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
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

func AccessLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)

	// Calls the handler
	reply, err := handler(ctx, req)

	var traceId, userId, userRole string
	if len(md["trace-id"]) > 0 {
		traceId = md["trace-id"][0]
	}
	if len(md["user-id"]) > 0 {
		userId = md["user-id"][0]
	}
	if len(md["user-role"]) > 0 {
		userRole = md["user-role"][0]
	}

	msg := fmt.Sprintf("Call:%v, traceId: %v, userId: %v, userRole: %v, time: %v", info.FullMethod, traceId, userId, userRole, time.Since(start))
	app.AccesLog(msg)

	return reply, err

}
