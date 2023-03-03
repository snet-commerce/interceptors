package errors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ErrToStatusMapperFunc represents function which maps error to grpc status
type ErrToStatusMapperFunc func(err error) *status.Status

// ServerUnaryInterceptor unary server interceptor for error to grpc status mapping
func ServerUnaryInterceptor(mapper ErrToStatusMapperFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		res, err := handler(ctx, req)
		if err == nil {
			return res, nil
		}

		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, mapper(err).Err()
	}
}
