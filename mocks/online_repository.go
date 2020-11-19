package mocks

import (
	"errors"
	"orders-srv-go/models"
)

// onlineRepository...
type onlineRepository struct {
	store *store
}

// GetByClient...
func (o onlineRepository) GetByClient(clientID int) (models.OnlineOrders, error) {
	o.store.t.Helper()
	if clientID == 0 {
		return nil, errors.New("client id is empty")
	}
	return models.OnlineOrders{models.OnlineOrder{ID: 1}}, nil
}
