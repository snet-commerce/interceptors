package errors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type StatusToErrMapper func(s *status.Status) error

func ClientUnaryInterceptor(mapper StatusToErrMapper) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, req, cc)
		s, ok := status.FromError(err)
		if !ok {
			return err
		}
		return mapper(s)
	}
}

//func ClientStreamInterceptor(mapper StatusToErrMapper) grpc.StreamClientInterceptor {
//	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
//		streamer(ctx, desc, cc, method, opts...)
//	}
//}
