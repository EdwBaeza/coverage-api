package infraestructure

import (
	"time"

	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	"github.com/stretchr/testify/mock"
)

// TODO: Improve mongo repository to reuse base logic (CRUD by generic types maybe)
type MockRepository struct {
	mock.Mock
}

func (repository *MockRepository) Create(purchase domain.Purchase) (domain.Purchase, error) {
	result := repository.Called(purchase)
	argument0 := result.Get(0)

	if argument0 == nil {
		return domain.Purchase{}, result.Error(1)
	}

	purchaseCreated := argument0.(domain.Purchase)
	return purchaseCreated, nil
}

func (repository *MockRepository) Find(id string) (domain.Purchase, error) {
	result := repository.Called(id)
	argument0 := result.Get(0)
	if argument0 == nil {
		return domain.Purchase{}, result.Error(1)
	}
	purchaseFound := argument0.(domain.Purchase)
	purchaseFound.ID = id

	return purchaseFound, nil
}

func (repository *MockRepository) List(filters map[string]interface{}) ([]domain.Purchase, error) {
	result := repository.Called()
	argument0 := result.Get(0)
	if argument0 == nil {
		return []domain.Purchase{}, result.Error(1)
	}
	purchasesFound := argument0.([]domain.Purchase)
	return purchasesFound, nil
}

func (repository *MockRepository) UpdateStatus(id string, status domain.PurchaseStatus) (domain.Purchase, error) {
	result := repository.Called(id, status)
	argument0 := result.Get(0)

	if result.Get(0) == nil {
		return domain.Purchase{}, result.Error(1)
	}

	purchaseUpdated := argument0.(domain.Purchase)
	purchaseUpdated.ID = id
	purchaseUpdated.UpdatedAt = time.Now()

	return purchaseUpdated, nil
}
