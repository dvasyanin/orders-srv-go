package models

import "time"

type OfflineOrders []OfflineOrder

type OfflineOrder struct {
	tableName      struct{} `pg:"offline_orders"`
	Id             int64    `pg:",pk"`
	StoreName      string   `pg:"store_name"`
	StoreId        int      `pg:"store_id"`
	CashboxId      int      `pg:"cashbox_id"`
	OrderId        string   `pg:"order_id"`
	RowReceipt     int      `pg:"row_receipt"`
	Date           time.Time
	Time           string `pg:"-"`
	Article        string
	Title          string
	Color          string
	Size           string
	Barcode        string
	Qty            int
	TotalGross     int `pg:"total_gross"`
	Discount       int
	Total          int
	ClientId       int `pg:"client_id"`
	Seller         string
	Family         string
	Operation      int
	Source         string
	BonusesWasted  int `pg:"bonuses_wasted"`
	BonusesAccrued int `pg:"bonuses_accrued"`
	Hash           string
	CreatedAt      time.Time `pg:"created_at"`
	UpdatedAt      time.Time `pg:"updated_at"`
}
