package usecase

import (
	"orders-srv-go/models"
	"time"
)

// OfflineOrders...
type OfflineOrders interface {
	ByClient(limit, offset, clientID int) (models.OfflineOrders, int, error)
	GetAll(limit, offset int, createDateFrom, createDateTo time.Time) (models.OfflineOrders, int, error)
}

// onlineOrders...
type OnlineOrders interface {
	ByClient(clientID int) (models.OnlineOrders, error)
}
