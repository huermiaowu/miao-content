package service

import (
	"content/internal/db"
	"context"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"

	"content/internal/svc"
	"content/pb"
)

type UpdateDynamicService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDynamicService(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDynamicService {
	return &UpdateDynamicService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *UpdateDynamicService) UpdateDynamic(in *pb.UpdateDynamicReq) (*pb.UpdateDynamicResp, error) {
	// todo: add your logic here and delete this line

	// todo: 借助client查询user，pet是否存在
	dynamic := db.Dynamic{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
		Title:   in.Title,
		Content: in.Content,
		UserId:  in.UserId,
		PetId:   in.PetId,
	}
	err := db.UpdateDynamic(s.ctx, &dynamic)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	dynamicInfo := pb.DynamicInfo{
		Id:      uint64(dynamic.ID),
		Title:   dynamic.Title,
		Content: dynamic.Content,
	}
	return &pb.UpdateDynamicResp{DynamicInfo: &dynamicInfo}, nil
}
