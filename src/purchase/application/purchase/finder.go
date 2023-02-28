package purchase

import (
	"errors"

	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	"github.com/edwbaeza/coverage-api/src/shared"
	applicationShared "github.com/edwbaeza/coverage-api/src/shared/application/validators"
	domainUser "github.com/edwbaeza/coverage-api/src/user/domain"
)

type Finder struct {
	repository domain.PurchaseRepository
}

func NewPurchaseFinder(repository domain.PurchaseRepository) *Finder {
	return &Finder{
		repository: repository,
	}
}

func (finder *Finder) Find(purchasId string, currentUser domainUser.User) (domain.Purchase, error) {
	purchase, err := finder.repository.Find(purchasId)

	if err != nil {
		return domain.Purchase{}, err
	}

	if !applicationShared.CanInteractWithPurchase(purchase, currentUser) {
		return domain.Purchase{}, shared.NewError(errors.New("Unauthorized action for this purchase"), "purchase.id")
	}

	return purchase, nil
}
