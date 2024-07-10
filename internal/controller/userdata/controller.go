package userdata

import (
	"context"
	"errors"
	"go/grpc/userservice/internal/repository"
	"go/grpc/userservice/internal/repository/memory/specification"
	"go/grpc/userservice/pkg/model"
)

var ErrNotFound = errors.New("not found")

type userdataRepository interface {
	GetUserByID(ctx context.Context, id int) (*model.UserData, error)
	GetUsersByIDs(ctx context.Context, ids []int) ([]*model.UserData, error)
	SearchUsers(ctx context.Context, specs ...specification.Specification) ([]*model.UserData, error)
}

type Controller struct {
	repo userdataRepository
}

func New(repo userdataRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetUserByID(ctx context.Context, id int) (*model.UserData, error) {
	res, err := c.repo.GetUserByID(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}

func (c *Controller) GetUsersByIDs(ctx context.Context, ids []int) ([]*model.UserData, error) {
	res, err := c.repo.GetUsersByIDs(ctx, ids)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}

func (c *Controller) SearchUsers(ctx context.Context, firstName, lastName, city string, married bool, weight float32) ([]*model.UserData, error) {
	specs := []specification.Specification{
		&specification.FirstNameSpecification{FirstName: firstName},
		&specification.LastNameSpecification{LastName: lastName},
		&specification.CitySpecification{City: city},
		&specification.MarriedSpecification{Married: married},
		&specification.WeightGreaterThanSpecification{Weight: weight},
	}
	res, err := c.repo.SearchUsers(ctx, specs...)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
