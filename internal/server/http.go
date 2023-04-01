package server

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/gorilla/handlers"
	v1 "kratos/api/realworld/v1"
	"kratos/internal/conf"
	"kratos/internal/pkg/middleware/auth"
	"kratos/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewSkipRoutersMatcher() selector.MatchFunc {
	skipRouters := make(map[string]struct{})
	skipRouters["/realworld.v1.Greeter/Login"] = struct{}{}
	skipRouters["/realworld.v1.Greeter/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}

}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, greeter *service.RealWorldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwtc.Token)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			func(h nethttp.Handler) nethttp.Handler {
				return nethttp.HandlerFunc(func(writer nethttp.ResponseWriter, request *nethttp.Request) {
					fmt.Println("route in")
					h.ServeHTTP(writer, request)
					fmt.Println("route out")
				})
			},
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
