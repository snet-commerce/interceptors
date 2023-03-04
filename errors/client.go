package errors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// StatusToErrMapper represents function which maps grpc status to error
type StatusToErrMapper func(s *status.Status) error

// ClientUnaryInterceptor unary client interceptor grpc status to error mapping
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
