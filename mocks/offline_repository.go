package mocks

import (
	"errors"
	"orders-srv-go/models"
	"time"
)

// offlineRepository...
type offlineRepository struct {
	store *store
}

// GetByClient...
func (o offlineRepository) GetByClient(limit, offset, clientID int) (models.OfflineOrders, int, error) {
	o.store.t.Helper()
	if clientID == 0 {
		return nil, 0, errors.New("client id is empty")
	}
	return models.OfflineOrders{models.OfflineOrder{Id: 1}}, 1, nil
}

// GetAll...
func (o offlineRepository) GetAll(limit, offset int, createDateFrom, createDateTo time.Time) (models.OfflineOrders, error) {
	o.store.t.Helper()
	var defaultDate time.Time
	if createDateFrom == defaultDate || createDateTo == defaultDate {
		return nil, errors.New("create date from or create date to is empty")
	}
	return models.OfflineOrders{models.OfflineOrder{Id: 1}}, nil
}

// CountOfUniqueOrders...
func (o offlineRepository) CountOfUniqueOrders(createDateFrom, createDateTo time.Time) (int, error) {
	o.store.t.Helper()
	var defaultDate time.Time
	if createDateFrom == defaultDate || createDateTo == defaultDate {
		return 0, errors.New("create date from or create date to is empty")
	}
	return 1, nil
}
