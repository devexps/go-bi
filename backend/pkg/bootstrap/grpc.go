package bootstrap

import (
	"context"
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/metadata"
	"github.com/devexps/go-micro/v2/middleware/recovery"
	"github.com/devexps/go-micro/v2/middleware/tracing"
	"github.com/devexps/go-micro/v2/registry"
	microGrpc "github.com/devexps/go-micro/v2/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

const defaultTimeout = 5 * time.Second

// CreateGrpcClient creates a GRPC client
func CreateGrpcClient(ctx context.Context, r registry.Discovery, serviceName string, timeoutDuration *durationpb.Duration) grpc.ClientConnInterface {
	timeout := defaultTimeout
	if timeoutDuration != nil {
		timeout = timeoutDuration.AsDuration()
	}

	endpoint := "discovery:///" + serviceName

	conn, err := microGrpc.DialInsecure(
		ctx,
		microGrpc.WithEndpoint(endpoint),
		microGrpc.WithDiscovery(r),
		microGrpc.WithTimeout(timeout),
		microGrpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
			tracing.Client(),
		),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}
	return conn
}

// CreateGrpcServer Create GRPC server
func CreateGrpcServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *microGrpc.Server {
	var opts []microGrpc.ServerOption

	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, metadata.Server())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, microGrpc.Middleware(ms...))

	if cfg.Server.Grpc.Network != "" {
		opts = append(opts, microGrpc.Network(cfg.Server.Grpc.Network))
	}
	if cfg.Server.Grpc.Addr != "" {
		opts = append(opts, microGrpc.Address(cfg.Server.Grpc.Addr))
	}
	if cfg.Server.Grpc.Timeout != nil {
		opts = append(opts, microGrpc.Timeout(cfg.Server.Grpc.Timeout.AsDuration()))
	}
	return microGrpc.NewServer(opts...)
}
