package middleware

import (
	"context"
	"go/grpc/userservice/gen"
	"go/grpc/userservice/internal/validation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateSearchUsers validates the input of SearchUsersRequest
func ValidateSearchUsers(ctx context.Context, req interface{}) error {
	searchReq, ok := req.(*gen.SearchUsersRequest)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "invalid request type")
	}

	if err := validation.ValidateStringWithSpace(searchReq.FirstName); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid first name: %v", err)
	}
	if err := validation.ValidateStringWithSpace(searchReq.LastName); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid last name: %v", err)
	}
	if err := validation.ValidateStringWithSpace(searchReq.City); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid city: %v", err)
	}
	if err := validation.ValidateHeight(searchReq.HeightGreaterThan); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid height: %v", err)
	}
	if err := validation.ValidatePhone(searchReq.Phone); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid phone: %v", err)
	}

	return nil
}
