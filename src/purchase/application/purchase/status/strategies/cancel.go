package strategies

import (
	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
)

// May be more later be needed to implement the strategy pattern by structs instead of functions
// for this propuse it's sufficient to use a function's structure

func ProcessCancel(id string, repository domain.PurchaseRepository) error {
	// TODO: Process cancel logic with refund workflow

	return nil
}
