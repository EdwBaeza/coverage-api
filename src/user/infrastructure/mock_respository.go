package infrastructure

import (
	"time"

	domain "github.com/edwbaeza/coverage-api/src/user/domain"
	"github.com/stretchr/testify/mock"
)

// TODO: Improve mongo repository to reuse base logic (CRUD by generic types maybe)
type MockRepository struct {
	mock.Mock
}

func (repository *MockRepository) Create(user domain.User) (domain.User, error) {
	result := repository.Called(user)
	argument0 := result.Get(0)

	if argument0 == nil {
		return domain.User{}, result.Error(1)
	}

	userCreated := argument0.(domain.User)
	return userCreated, nil
}

func (repository *MockRepository) Find(id string) (domain.User, error) {
	result := repository.Called(id)
	argument0 := result.Get(0)
	if argument0 == nil {
		return domain.User{}, result.Error(1)
	}
	userFound := argument0.(domain.User)
	userFound.ID = id

	return userFound, nil
}

func (repository *MockRepository) FindByEmail(email string) (domain.User, error) {
	result := repository.Called(email)
	argument0 := result.Get(0)
	argument1 := result.Get(1)
	if result.Get(0) == nil || argument1 != nil {
		return domain.User{}, result.Error(1)
	}

	userFound := argument0.(domain.User)
	return userFound, nil
}

func (repository *MockRepository) Update(id string, user domain.User) (domain.User, error) {
	result := repository.Called(id, user)
	argument0 := result.Get(0)

	if result.Get(0) == nil {
		return domain.User{}, result.Error(1)
	}

	userUpdated := argument0.(domain.User)
	userUpdated.ID = id
	userUpdated.UpdatedAt = time.Now()

	return userUpdated, nil
}
