package repository

import (
	"orders-srv-go/models"
	"time"
)

// OnlineOrdersRepository...
type OnlineOrdersRepository interface {
	GetByClient(clientID int) (models.OnlineOrders, error)
}

// OfflineOrdersRepository...
type OfflineOrdersRepository interface {
	GetByClient(limit, offset, clientID int) (models.OfflineOrders, int, error)
	GetAll(limit, offset int, createDateFrom, createDateTo time.Time) (models.OfflineOrders, error)
	CountOfUniqueOrders(createDateFrom, createDateTo time.Time) (int, error)
}
