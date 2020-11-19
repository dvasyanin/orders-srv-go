package repository

import (
	"fmt"
	"orders-srv-go/models"
)

// onlineRepository...
type onlineRepository struct {
	store *store
}

// GetByClient...
func (o onlineRepository) GetByClient(clientID int) (models.OnlineOrders, error) {
	return models.OnlineOrders{}, fmt.Errorf("impliment me")
}
