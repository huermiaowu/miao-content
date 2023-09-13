package router

import (
	"context"
	"github.com/huerni/gmitex/pkg/gw"
)

type ContentRouter struct {
	Paths     []string
	AuthPaths []string
}

func NewContentRouter() *ContentRouter {
	paths := make([]string, 0)
	authPaths := make([]string, 0)

	paths = append(paths, "/api/v1/dynamic")

	return &ContentRouter{Paths: paths, AuthPaths: authPaths}
}

func (ur *ContentRouter) AddRouters(ctx context.Context, client gw.GatewayClient, info any) error {
	return client.AddRoute(ctx, info)
}
