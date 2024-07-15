package memory

import (
	"context"
	"fmt"
	"go/grpc/userservice/internal/repository"
	"go/grpc/userservice/internal/repository/memory/specification"
	"go/grpc/userservice/pkg/model"
	"math/rand"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[int]*model.UserData
}

func New() *Repository {
	rep := &Repository{
		data: map[int]*model.UserData{},
	}

	rep.GenerateRandomUserData(50, 123)
	return rep
}

// GetUserByID retrieves a user by user id.
func (r *Repository) GetUserByID(_ context.Context, id int) (*model.UserData, error) {
	r.RLock()
	defer r.RUnlock()
	user, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return user, nil
}

// GetUsersByIDs returns a list of users based on the IDs provided.
func (r *Repository) GetUsersByIDs(_ context.Context, ids []int) ([]*model.UserData, error) {
	r.RLock()
	defer r.RUnlock()
	var users []*model.UserData
	for _, id := range ids {
		if user, ok := r.data[id]; ok {
			users = append(users, user)
		}
	}

	if len(users) == 0 {
		return nil, repository.ErrNotFound
	}
	return users, nil
}

// SearchUsers returns a list of users based on the provided specifications.
func (r *Repository) SearchUsers(ctx context.Context, specs ...specification.Specification) ([]*model.UserData, error) {
	r.RLock()
	defer r.RUnlock()
	var users []*model.UserData
	andSpec := &specification.AndSpecification{Specs: specs}
	for _, user := range r.data {
		if andSpec.IsSatisfiedBy(ctx, user) {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		return nil, repository.ErrNotFound
	}
	return users, nil
}

// GenerateRandomUserData generates n amount of random user data.
func (r *Repository) GenerateRandomUserData(n int, seed int64) {
	rnd := rand.New(rand.NewSource(seed))

	firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Karen", "Leo", "Mia"}
	lastNames := []string{"Doe", "Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Martinez", "Miller", "Davis", "Wilson", "Taylor", "Clark", "Moore", "Hall"}
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose", "Austin", "Jacksonville", "Fort Worth", "Columbus", "San Francisco"}

	for i := 0; i < n; i++ {
		id := i + 1
		firstName := firstNames[rnd.Intn(len(firstNames))]
		lastName := lastNames[rnd.Intn(len(lastNames))]
		city := cities[rnd.Intn(len(cities))]
		phone := fmt.Sprintf("+1-555-%04d", rnd.Intn(10000))
		height := rnd.Float32()*2 + 1.5 // random height between 1.5 and 3.5 meters
		married := rnd.Intn(2) == 0     // random boolean

		r.data[id] = &model.UserData{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			City:      city,
			Phone:     phone,
			Height:    height,
			Married:   married,
		}
	}
}
