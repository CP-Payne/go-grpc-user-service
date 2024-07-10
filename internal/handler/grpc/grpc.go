package grpc

import (
	"context"
	"errors"
	"go/grpc/userservice/gen"
	"go/grpc/userservice/internal/controller/userdata"
	"go/grpc/userservice/pkg/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedUserServiceServer
	ctrl *userdata.Controller
}

func New(ctrl *userdata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetUserByID(ctx context.Context, req *gen.GetUserByIDRequest) (*gen.GetUserResponse, error) {
	if req == nil || req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	user, err := h.ctrl.GetUserByID(ctx, int(req.Id))
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetUserResponse{User: model.UserdataToProto(user)}, nil
}

// TODO: What if 5 ids are provided and 3 does not exist?
// Add a field that specifies which IDs was not found
// Perform validation if incorrect values are provided
// Maybe return empty list if no users is found in place of an error
func (h *Handler) GetUsersByIDs(ctx context.Context, req *gen.GetUsersByIDsRequest) (*gen.GetUsersResponse, error) {
	if req == nil || len(req.Ids) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or no IDs")
	}
	users, err := h.ctrl.GetUsersByIDs(ctx, model.IDsInt32ToInt(req.Ids))
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetUsersResponse{Users: model.UsersdataToProto(users)}, nil
}

func (h *Handler) SearchUsers(ctx context.Context, req *gen.SearchUsersRequest) (*gen.GetUsersResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	users, err := h.ctrl.SearchUsers(ctx, req.FirstName, req.LastName, req.City, req.Married, req.HeightGreaterThan)
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetUsersResponse{Users: model.UsersdataToProto(users)}, nil
}
