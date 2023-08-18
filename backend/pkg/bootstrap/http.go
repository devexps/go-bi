package bootstrap

import (
	"github.com/devexps/go-bi/api/gen/go/common/conf"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/metadata"
	"github.com/devexps/go-micro/v2/middleware/recovery"
	"github.com/devexps/go-micro/v2/middleware/tracing"
	microHttp "github.com/devexps/go-micro/v2/transport/http"
	"github.com/gorilla/handlers"
)

// CreateHttpServer creates a HTTP server
func CreateHttpServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *microHttp.Server {
	var opts = []microHttp.ServerOption{
		microHttp.Filter(handlers.CORS(
			handlers.AllowedHeaders(cfg.Server.Http.Headers),
			handlers.AllowedMethods(cfg.Server.Http.Methods),
			handlers.AllowedOrigins(cfg.Server.Http.Origins),
		)),
	}

	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, metadata.Server())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, microHttp.Middleware(ms...))

	if cfg.Server.Http.Network != "" {
		opts = append(opts, microHttp.Network(cfg.Server.Http.Network))
	}
	if cfg.Server.Http.Addr != "" {
		opts = append(opts, microHttp.Address(cfg.Server.Http.Addr))
	}
	if cfg.Server.Http.Timeout != nil {
		opts = append(opts, microHttp.Timeout(cfg.Server.Http.Timeout.AsDuration()))
	}
	return microHttp.NewServer(opts...)
}
