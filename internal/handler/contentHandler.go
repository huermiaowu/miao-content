// no edit

package handler

import (
	"context"

	"github.com/huermiaowu/miao-content/internal/service"
	"github.com/huermiaowu/miao-content/internal/svc"
	"github.com/huermiaowu/miao-content/pb"
)

type ContentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedContentServer
}

func NewContentServer(svcCtx *svc.ServiceContext) *ContentServer {
	return &ContentServer{
		svcCtx: svcCtx,
	}
}

// Ping /api/v1/dynamic/ping
func (s *ContentServer) Ping(ctx context.Context, in *pb.PingReq) (*pb.PingResp, error) {
	res, err := service.NewPingService(ctx, s.svcCtx).Ping(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateDynamic /api/v1/dynamic/create
func (s *ContentServer) CreateDynamic(ctx context.Context, in *pb.CreateDynamicReq) (*pb.CreateDynamicResp, error) {
	res, err := service.NewCreateDynamicService(ctx, s.svcCtx).CreateDynamic(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateDynamic /api/v1/dynamic/update
func (s *ContentServer) UpdateDynamic(ctx context.Context, in *pb.UpdateDynamicReq) (*pb.UpdateDynamicResp, error) {
	res, err := service.NewUpdateDynamicService(ctx, s.svcCtx).UpdateDynamic(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetDynamic /api/v1/dynamic
func (s *ContentServer) GetDynamic(ctx context.Context, in *pb.GetDynamicReq) (*pb.GetDynamicResp, error) {
	res, err := service.NewGetDynamicService(ctx, s.svcCtx).GetDynamic(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
