package main

import (
	"content/internal/app"
	"content/internal/config"
	"content/internal/handler"
	"content/internal/svc"
	content "content/pb"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/huerni/gmitex/pkg/http/handlers"
	"github.com/huerni/gmitex/pkg/http/response"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c, err := config.InitConfig("etc/cfg.toml")
	if err != nil {
		panic(err)
	}

	g := app.NewGmServer(c, content.RegisterContentHandlerFromEndpoint, func(server *grpc.Server) {
		ctx := svc.NewServiceContext(c)
		srv := handler.NewContentServer(ctx)
		content.RegisterContentServer(server, srv)
	})

	g.HttpServer.AddMuxOp(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &response.CustomMarshaler{
				M: &runtime.JSONPb{
					MarshalOptions:   protojson.MarshalOptions{},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}}}),
		runtime.WithErrorHandler(handlers.ErrorHandler),
	)

	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
