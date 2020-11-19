package service

import (
	"context"
	pb "github.com/12Storeez/esb-protobufs/go/orders"
	"github.com/micro/go-micro/errors"
	"github.com/zhs/loggr"
	"go.uber.org/zap"
	"orders-srv-go/internal"
	"orders-srv-go/usecase"
	"time"
)

// offlineOrders...
type offlineOrders struct {
	logger        *zap.SugaredLogger
	offlineOrders usecase.OfflineOrders
}

// NewOfflineOrders...
func NewOfflineOrders(logger *zap.SugaredLogger, offlineOrdersService usecase.OfflineOrders) *offlineOrders {
	return &offlineOrders{
		logger:        logger,
		offlineOrders: offlineOrdersService,
	}
}

// ByClient...
func (o offlineOrders) ByClient(ctx context.Context, params *pb.ParamsOfflineByClient, response *pb.ResponseOffline) error {
	log := o.logger.With(
		"context", "Offline.ByClient",
		"request_id", internal.GetMetadataField(ctx, internal.ConstRequestID),
		"request", loggr.Marshal(params))
	log.Info("[request]")

	orders, count, err := o.offlineOrders.ByClient(int(params.GetLimit()), int(params.GetOffset()), int(params.GetClientId()))
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("Offline.ByClient:", "%v", err)
	}

	response.Orders = make([]*pb.OfflineOrderPosition, len(orders))
	for i, order := range orders {
		response.Orders[i] = &pb.OfflineOrderPosition{
			StoreName:      order.StoreName,
			StoreId:        int32(order.StoreId),
			CashboxId:      int32(order.CashboxId),
			OrderId:        order.OrderId,
			RowReceipt:     int32(order.RowReceipt),
			Date:           timestampToString(order.Date),
			Time:           order.Time,
			Article:        order.Article,
			Title:          order.Title,
			Color:          order.Color,
			Size:           order.Size,
			Barcode:        order.Barcode,
			Qty:            int32(order.Qty),
			TotalGross:     int32(order.TotalGross),
			Discount:       int32(order.Discount),
			Total:          int32(order.Total),
			ClientId:       int32(order.ClientId),
			Seller:         order.Seller,
			Family:         order.Family,
			Operation:      int32(order.Operation),
			Source:         order.Source,
			BonusesWasted:  int32(order.BonusesWasted),
			BonusesAccrued: int32(order.BonusesAccrued),
		}
	}
	response.Total = int32(count)
	log.Infow("[response]", "response", loggr.Marshal(response))

	return nil
}

// GetAll...
func (o offlineOrders) GetAll(ctx context.Context, params *pb.ParamsGetAll, response *pb.ResponseOffline) error {
	log := o.logger.With(
		"context", "Offline.GetAll",
		"request_id", internal.GetMetadataField(ctx, internal.ConstRequestID),
		"request", loggr.Marshal(params))
	log.Info("[request]")

	createdDateFrom := time.Unix(int64(params.GetCreateDateFrom()), 0)
	createdDateTo := time.Unix(int64(params.GetCreateDateTo()), 0)

	orders, count, err := o.offlineOrders.GetAll(int(params.GetLimit()), int(params.GetOffset()), createdDateFrom, createdDateTo)
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("Offline.GetAll:", "%v", err)
	}

	response.Orders = make([]*pb.OfflineOrderPosition, len(orders))
	for i, order := range orders {
		response.Orders[i] = &pb.OfflineOrderPosition{
			StoreName:      order.StoreName,
			StoreId:        int32(order.StoreId),
			CashboxId:      int32(order.CashboxId),
			OrderId:        order.OrderId,
			RowReceipt:     int32(order.RowReceipt),
			Date:           timestampToString(order.Date),
			Time:           order.Time,
			Article:        order.Article,
			Title:          order.Title,
			Color:          order.Color,
			Size:           order.Size,
			Barcode:        order.Barcode,
			Qty:            int32(order.Qty),
			TotalGross:     int32(order.TotalGross),
			Discount:       int32(order.Discount),
			Total:          int32(order.Total),
			ClientId:       int32(order.ClientId),
			Seller:         order.Seller,
			Family:         order.Family,
			Operation:      int32(order.Operation),
			Source:         order.Source,
			BonusesWasted:  int32(order.BonusesWasted),
			BonusesAccrued: int32(order.BonusesAccrued),
		}
	}
	response.Total = int32(count)
	log.Infow("[response]", "response", loggr.Marshal(response))

	return nil
}

// timestampToString...
func timestampToString(ts time.Time) string {
	return ts.Format("2006-01-02 15:04:05")
}
