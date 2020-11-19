package service

import (
	"context"
	pb "github.com/12Storeez/esb-protobufs/go/orders"
	"github.com/micro/go-micro/errors"
	"github.com/zhs/loggr"
	"go.uber.org/zap"
	"orders-srv-go/internal"
	"orders-srv-go/usecase"
)

// onlineOrders...
type onlineOrders struct {
	logger       *zap.SugaredLogger
	onlineOrders usecase.OnlineOrders
}

// NewOnlineOrders...
func NewOnlineOrders(logger *zap.SugaredLogger, onlineOrdersUsecase usecase.OnlineOrders) *onlineOrders {
	return &onlineOrders{
		logger:       logger,
		onlineOrders: onlineOrdersUsecase,
	}
}

// ByClient...
func (o onlineOrders) ByClient(ctx context.Context, params *pb.ParamsOnlineByClient, response *pb.ResponseOnlineByClient) error {
	log := o.logger.With(
		"context", "Online.ByClient",
		"request_id", internal.GetMetadataField(ctx, internal.ConstRequestID),
		"request", loggr.Marshal(params))
	log.Info("[request]")

	err := errors.InternalServerError("Online.ByClient", "%s", "implement me")
	log.Error(err)
	return err
}
