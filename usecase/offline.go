package usecase

import (
	"orders-srv-go/models"
	"orders-srv-go/repository"
	"sync"
	"time"
)

// NewOfflineOrders...
func NewOfflineOrders(store repository.OrderStore) OfflineOrders {
	return &offlineOrders{
		store: store,
	}
}

// offlineOrders...
type offlineOrders struct {
	store repository.OrderStore
}

// GetAll...
func (u offlineOrders) GetAll(limit, offset int, createDateFrom, createDateTo time.Time) (models.OfflineOrders, int, error) {
	wg := &sync.WaitGroup{}
	errs := make(chan error, 1)

	var count int
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		count, err = u.store.Offline().CountOfUniqueOrders(createDateFrom, createDateTo)
		errs <- err
	}()

	var orders models.OfflineOrders
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		orders, err = u.store.Offline().GetAll(limit, offset, createDateFrom, createDateTo)
		errs <- err
	}()

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return nil, 0, err
		}
	}

	return orders, count, nil
}

// ByClient...
func (u offlineOrders) ByClient(limit, offset, clientID int) (models.OfflineOrders, int, error) {
	return u.store.Offline().GetByClient(limit, offset, clientID)
}
