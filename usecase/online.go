package usecase

import (
	"orders-srv-go/models"
	"orders-srv-go/repository"
)

// NewOnlineOrders...
func NewOnlineOrders(store repository.OrderStore) OnlineOrders {
	return &onlineOrders{
		store: store,
	}
}

// onlineOrders...
type onlineOrders struct {
	store repository.OrderStore
}

// ByClient...
func (u onlineOrders) ByClient(clientID int) (models.OnlineOrders, error) {
	return u.store.Online().GetByClient(clientID)
}
