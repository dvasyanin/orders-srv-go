package main

import (
	pb "github.com/12Storeez/esb-protobufs/go/orders"
	"github.com/heptiolabs/healthcheck"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/zhs/loggr"
	"net/http"
	"orders-srv-go/config"
	k8s_health "orders-srv-go/pkg/k8s-health"
	"orders-srv-go/pkg/postgres"
	"orders-srv-go/repository"
	"orders-srv-go/service"
	"orders-srv-go/usecase"
	"time"
)

func main() {
	cfg := config.New()
	logger := loggr.New("@version", cfg.App.Version)
	pgdb := postgres.NewPostgres(cfg.Postgres.Addr, cfg.Postgres.Database, cfg.Postgres.User, cfg.Postgres.Password)
	store := repository.NewStore(pgdb)

	//health
	health := healthcheck.NewHandler()
	health.AddLivenessCheck("service check", healthcheck.TCPDialCheck("localhost:50051", 5*time.Second))
	health.AddReadinessCheck("pg check", k8s_health.PostgresPingCheck(pgdb, 5*time.Second))
	go http.ListenAndServe("0.0.0.0:8086", health)

	// MICRO gRPC SERVER
	srv := micro.NewService(
		micro.Name("orders"),
		micro.Address(cfg.App.Port),
	)

	srv.Init()
	_ = pb.RegisterOfflineHandler(srv.Server(), service.NewOfflineOrders(logger, usecase.NewOfflineOrders(store)))
	_ = pb.RegisterOnlineHandler(srv.Server(), service.NewOnlineOrders(logger, usecase.NewOnlineOrders(store)))

	if err := srv.Run(); err != nil {
		log.Fatalf("can't run server: %v", err)
	}
}
