package mw

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(ctx, "[interceptor.Logging] method: %s; metadata: %v", info.FullMethod, md)
	}

	res, err := handler(ctx, req)
	if err != nil {
		log.Println(ctx, "[interceptor.Logging] method: %s; error: %s", info.FullMethod, err.Error())
		return
	}

	respReq, _ := protojson.Marshal((res).(proto.Message))
	log.Println(ctx, "[interceptor.Logging] method: %s; response: %s", info.FullMethod, string(respReq))
	return res, nil
}
