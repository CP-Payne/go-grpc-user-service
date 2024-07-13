package userdata

import (
	"context"
	"errors"
	"go/grpc/userservice/internal/repository"
	"go/grpc/userservice/internal/repository/memory/specification"
	"go/grpc/userservice/pkg/model"
)

// ErrNotFound is returned when a requested record is not found.
var ErrNotFound = errors.New("not found")

type userdataRepository interface {
	GetUserByID(ctx context.Context, id int) (*model.UserData, error)
	GetUsersByIDs(ctx context.Context, ids []int) ([]*model.UserData, error)
	SearchUsers(ctx context.Context, specs ...specification.Specification) ([]*model.UserData, error)
}

// Controller defines a userdata service controller.
type Controller struct {
	repo userdataRepository
}

// New creates a userdata service controller.
func New(repo userdataRepository) *Controller {
	return &Controller{repo}
}

// GetUserByID returns user data by id.
func (c *Controller) GetUserByID(ctx context.Context, id int) (*model.UserData, error) {
	res, err := c.repo.GetUserByID(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}

// GetUsersByIDs returns user data for the provided ids.
func (c *Controller) GetUsersByIDs(ctx context.Context, ids []int) ([]*model.UserData, error) {
	res, err := c.repo.GetUsersByIDs(ctx, ids)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}

// SearchUsers returns user data based on the provided specifications.
func (c *Controller) SearchUsers(ctx context.Context, firstName, lastName, city, phone string, married *bool, weight float32) ([]*model.UserData, error) {
	specs := []specification.Specification{
		&specification.FirstNameSpecification{FirstName: firstName},
		&specification.LastNameSpecification{LastName: lastName},
		&specification.CitySpecification{City: city},
		&specification.WeightGreaterThanSpecification{Weight: weight},
		&specification.PhoneSpecification{Phone: phone},
	}

	if married != nil {
		specs = append(specs, &specification.MarriedSpecification{Married: *married, IsSet: true})
	} else {
		specs = append(specs, &specification.MarriedSpecification{IsSet: false})
	}

	res, err := c.repo.SearchUsers(ctx, specs...)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
