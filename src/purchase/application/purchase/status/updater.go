package statu

import (
	"errors"
	"fmt"

	applicationPurchase "github.com/edwbaeza/coverage-api/src/purchase/application/purchase"
	"github.com/edwbaeza/coverage-api/src/purchase/application/purchase/status/strategies"
	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	"github.com/edwbaeza/coverage-api/src/shared"
	domainUser "github.com/edwbaeza/coverage-api/src/user/domain"
)

// TODO: Improve this to use a state machine if it's needed
var (
	TRANSITIONS = map[domain.PurchaseStatus][]domain.PurchaseStatus{
		domain.Created: {
			domain.Created, // TODO: Mokey patch to allow the same status
		},
		domain.Collected: {
			domain.Created,
		},
		domain.InWarehouse: {
			domain.Collected,
		},
		domain.OnRoute: {
			domain.InWarehouse,
		},
		domain.Delivered: {
			domain.OnRoute,
		},
		domain.Canceled: {
			domain.Created,
			domain.Collected,
			domain.InWarehouse,
		},
	}
)

type StatusUpdater struct {
	repository domain.PurchaseRepository
	finder     applicationPurchase.Finder
}

func NewPurchaseStatusUpdater(repository domain.PurchaseRepository) *StatusUpdater {
	return &StatusUpdater{
		repository: repository,
		finder:     *applicationPurchase.NewPurchaseFinder(repository),
	}
}

func (updater *StatusUpdater) Update(id string, status int, currentUser domainUser.User) (domain.Purchase, error) {
	purchase, err := updater.finder.Find(id, currentUser)
	if err != nil {
		return domain.Purchase{}, err
	}

	statusConverted := domain.PurchaseStatus(status)
	if !updater.CanUpdate(purchase, statusConverted) {
		return domain.Purchase{},
			shared.NewError(errors.New(fmt.Sprintf("Invalid transition from %d to %d", purchase.Status, status)), "purchase.status")
	}

	purchase, err = updater.repository.UpdateStatus(id, statusConverted)

	if err != nil {
		return domain.Purchase{}, err
	}

	err = updater.ProcessStrategyByStatus(id, statusConverted)

	if err != nil {
		return domain.Purchase{}, err
	}

	return updater.repository.Find(id)
}

func (updater *StatusUpdater) CanUpdate(purchase domain.Purchase, status domain.PurchaseStatus) bool {
	allowedStatuses := TRANSITIONS[status]
	for _, allowedStatus := range allowedStatuses {
		if allowedStatus == purchase.Status {
			return true
		}
	}

	return false
}

func (updater *StatusUpdater) ProcessStrategyByStatus(id string, status domain.PurchaseStatus) error {
	switch status {
	case domain.Canceled:
		return strategies.ProcessCancel(id, updater.repository)
	}

	// It mean that the status strategy is not implemented yet
	return nil
}
