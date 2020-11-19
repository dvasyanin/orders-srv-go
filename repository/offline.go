package repository

import (
	"fmt"
	"orders-srv-go/models"
	"time"
)

const (
	layout        = "2006-01-02 15:04:05"
	limitDefault  = 1000
	limitMin      = 1
	offsetDefault = 0
)

// offlineRepository...
type offlineRepository struct {
	store *store
}

// GetAll...
func (o offlineRepository) GetAll(limit, offset int, createDateFrom, createDateTo time.Time) (models.OfflineOrders, error) {
	// set default limit
	if limit < 1 || limit > 1000 {
		limit = 1000
	}

	// set default offset
	if offset < 0 {
		offset = 0
	}

	queryResult := fmt.Sprintf(`SELECT *
		FROM offline_orders
		WHERE order_id IN (
    		SELECT order_id
    		FROM (
             	SELECT DISTINCT ON (order_id, store_id, date) *
             	FROM offline_orders
             	where created_at > '%s' and created_at < '%s'
         	) AS res
    	LIMIT %d offset %d);`, createDateFrom.Format(layout), createDateTo.Format(layout), limit, offset)

	var result models.OfflineOrders
	if _, err := o.store.db.Query(&result, queryResult); err != nil {
		return nil, err
	}

	return result, nil
}

// CountOfUniqueOrders...
func (o offlineRepository) CountOfUniqueOrders(createDateFrom, createDateTo time.Time) (int, error) {
	queryCount := fmt.Sprintf(`select count(*)
		from (select distinct on (order_id, store_id, date) *
		from offline_orders
      	where created_at > '%s' and created_at < '%s') as res`, createDateFrom.Format(layout), createDateTo.Format(layout))

	var count struct {
		Count int `sql:"count"`
	}
	if _, err := o.store.db.Query(&count, queryCount); err != nil {
		return 0, err
	}

	return count.Count, nil
}

// GetByClient...
func (o offlineRepository) GetByClient(limit, offset, clientId int) (models.OfflineOrders, int, error) {
	// set default limit
	if limit < limitMin || limit > limitDefault {
		limit = limitDefault
	}

	// set default offset
	if offset < offsetDefault {
		offset = offsetDefault
	}

	var result models.OfflineOrders
	count, err := o.store.db.Model(&result).
		Where("client_id = ?", clientId).
		Limit(limit).
		Offset(offset).
		SelectAndCount()
	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
}
