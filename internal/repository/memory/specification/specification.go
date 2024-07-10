package specification

import (
	"context"
	"go/grpc/userservice/pkg/model"
)

type Specification interface {
	IsSatisfiedBy(ctx context.Context, user *model.UserData) bool
}

type AndSpecification struct {
	Specs []Specification
}

func (s *AndSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	for _, spec := range s.Specs {
		if !spec.IsSatisfiedBy(ctx, user) {
			return false
		}
	}
	return true
}
