package grpc

import (
	"context"
	"errors"
	"go/grpc/userservice/gen"
	"go/grpc/userservice/internal/controller/userdata"
	"go/grpc/userservice/internal/decorator"
	"go/grpc/userservice/internal/validation"
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
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if err := validation.ValidateID(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID: %v", err)
	}

	user, err := h.ctrl.GetUserByID(ctx, int(req.Id))
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetUserResponse{User: model.UserdataToProto(user)}, nil
}

func (h *Handler) GetUsersByIDs(ctx context.Context, req *gen.GetUsersByIDsRequest) (*gen.GetUsersResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if err := validation.ValidateIDs(req.Ids); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid IDs: %v", err)
	}
	users, err := h.ctrl.GetUsersByIDs(ctx, model.IDsInt32ToInt(req.Ids))
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "no users found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	foundIDs := make(map[int32]bool)
	for _, user := range users {
		foundIDs[int32(user.ID)] = true
	}

	var notFoundIDs []int32
	for _, id := range req.Ids {
		if !foundIDs[id] {
			notFoundIDs = append(notFoundIDs, id)
		}
	}
	return &gen.GetUsersResponse{Users: model.UsersdataToProto(users), NotFoundIds: notFoundIDs}, nil
}

func (h *Handler) SearchUsers(ctx context.Context, req *gen.SearchUsersRequest) (*gen.GetUsersResponse, error) {
	decoratedHandler := decorator.ValidateSearchUsers(h.searchUsersCore)
	resp, err := decoratedHandler(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*gen.GetUsersResponse), nil
}

func (h *Handler) searchUsersCore(ctx context.Context, req interface{}) (interface{}, error) {
	searchReq := req.(*gen.SearchUsersRequest)
	var isMarried *bool
	if searchReq.Married != nil {
		isMarried = &searchReq.GetMarried().Value
	} else {
		isMarried = nil
	}
	users, err := h.ctrl.SearchUsers(ctx, searchReq.FirstName, searchReq.LastName, searchReq.City, isMarried, searchReq.HeightGreaterThan)
	if err != nil && errors.Is(err, userdata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "no users found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetUsersResponse{Users: model.UsersdataToProto(users)}, nil
}
