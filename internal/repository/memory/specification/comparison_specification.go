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

type WeightGreaterThanSpecification struct {
	Weight float32
}

func (s *WeightGreaterThanSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.Weight == 0 || user.Height > s.Weight
}

type MarriedSpecification struct {
	Married bool
	IsSet   bool // Indicates whether the Married value has been set
}

func (s *MarriedSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	if !s.IsSet {
		return true
	}
	return user.Married == s.Married
}

type PhoneSpecification struct {
	Phone string
}

func (s *PhoneSpecification) IsSatisfiedBy(ctx context.Context, user *model.UserData) bool {
	return s.Phone == "" || user.Phone == s.Phone
}
