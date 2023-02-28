package validators

import (
	doaminPurchase "github.com/edwbaeza/coverage-api/src/purchase/domain"
	domainUser "github.com/edwbaeza/coverage-api/src/user/domain"
)

func CanInteractWithPurchase(purchase doaminPurchase.Purchase, currentUser domainUser.User) bool {
	if currentUser.IsSuperAdmin() {
		return true
	}

	return purchase.UserId == currentUser.ID
}
