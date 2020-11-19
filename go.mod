module "orders-srv-go"

go 1.15

require (
	github.com/12Storeez/esb-protobufs v0.3.26
	github.com/go-pg/pg/v10 v10.6.2
	github.com/golang/protobuf v1.4.3
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/subosito/gotenv v1.2.0
	github.com/zhs/loggr v0.0.4
	go.uber.org/zap v1.14.1
	google.golang.org/grpc v1.27.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0