package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	md "github.com/ariefzainuri96/go-api-ecommerce-api-gateway/cmd/api/middleware"
)

func MetadataInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		// Extract request-id from incoming HTTP ctx
		reqID := ctx.Value(md.CtxRequestID)
		if reqID == nil {
			reqID = uuid.New().String()
		}

		md := metadata.Pairs(
			"x-correlation-id", fmt.Sprintf("%v", reqID),
			"source", "api-gateway",
		)

		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}