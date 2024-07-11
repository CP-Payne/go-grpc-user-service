package specification

import (
	"context"
	"go/grpc/userservice/pkg/model"
)

type FirstNameSpecification struct {
	FirstName string
}

func (s *FirstNameSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.FirstName == "" || user.FirstName == s.FirstName
}

type LastNameSpecification struct {
	LastName string
}

func (s *LastNameSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.LastName == "" || user.LastName == s.LastName
}

type CitySpecification struct {
	City string
}

func (s *CitySpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.City == "" || user.City == s.City
}

type MarriedSpecification struct {
	Married bool
}

func (s *MarriedSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {

	return user.Married == s.Married
}

type WeightGreaterThanSpecification struct {
	Weight float32
}

func (s *WeightGreaterThanSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.Weight == 0 || user.Height > s.Weight
}
