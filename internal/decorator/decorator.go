package decorator

import (
	"context"
	"go/grpc/userservice/internal/middleware"
)

type HandlerFunc func(ctx context.Context, req interface{}) (interface{}, error)

// ValidateSearchUsers decorates the request with validation rules.
func ValidateSearchUsers(next HandlerFunc) HandlerFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if err := middleware.ValidateSearchUsers(ctx, req); err != nil {
			return nil, err
		}
		return next(ctx, req)
	}
}
