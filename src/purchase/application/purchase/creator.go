package purchase

import (
	"errors"
	"time"

	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	"github.com/edwbaeza/coverage-api/src/shared"
	sharedDomain "github.com/edwbaeza/coverage-api/src/shared/domain"
	utils "github.com/edwbaeza/coverage-api/src/utils"
)

type Creator struct {
	repository domain.PurchaseRepository
}

func NewPurchaseCreator(repository domain.PurchaseRepository) *Creator {
	return &Creator{
		repository: repository,
	}
}

func (creator *Creator) Create(purchase domain.Purchase) (domain.Purchase, error) {
	// TODO: Validate coverage for addresses
	// TODO: Validate nested relationships (fk)

	if !creator.ValidateProductSizes(*purchase.Items) {
		return domain.Purchase{}, shared.NewError(errors.New("This purchase does not include the standard service"), "purchase.items")
	}

	totalPrice := utils.Sum(*purchase.Items, func(item domain.PurchaseItem) float64 {
		return float64(item.Units) * item.UnitPrice
	})

	purchase.Status = domain.Created
	purchase.TotalPrice = totalPrice
	purchase.BaseModel = sharedDomain.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return creator.repository.Create(purchase)
}

func (creator *Creator) ValidateProductSizes(items []domain.PurchaseItem) bool {
	sizes := map[int]int32{
		1: 5,
		2: 15,
		3: 25,
	}

	sizeAccumulator := int32(0)

	for _, item := range items {
		size := item.ProductSize
		units := item.Units
		sizeAccumulator = sizeAccumulator + (sizes[int(size)] * units)
	}

	return sizeAccumulator <= 25
}
