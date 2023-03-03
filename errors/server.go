package errors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ErrToStatusMapperFunc func(err error) *status.Status

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

func ServerStreamInterceptor(mapper ErrToStatusMapperFunc) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		if err == nil {
			return nil
		}

		if _, ok := status.FromError(err); ok {
			return err
		}

		return mapper(err).Err()
	}
}
