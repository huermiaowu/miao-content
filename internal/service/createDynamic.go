package service

import (
	"content/internal/db"
	"context"
	"github.com/huerni/gmitex/pkg/errno"
	"gorm.io/gorm"

	"content/internal/svc"
	"content/pb"
)

type CreateDynamicService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDynamicService(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDynamicService {
	return &CreateDynamicService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *CreateDynamicService) CreateDynamic(in *pb.CreateDynamicReq) (*pb.CreateDynamicResp, error) {
	// todo: add your logic here and delete this line

	// todo: 借助client查询user，pet是否存在
	dynamic := db.Dynamic{
		Model:   gorm.Model{},
		Title:   in.Title,
		Content: in.Content,
		UserId:  in.UserId,
		PetId:   in.PetId,
	}
	err := db.CreateDynamic(s.ctx, &dynamic)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	dynamicInfo := pb.DynamicInfo{
		Id:      uint64(dynamic.ID),
		Title:   dynamic.Title,
		Content: dynamic.Content,
	}
	return &pb.CreateDynamicResp{DynamicInfo: &dynamicInfo}, nil
}
