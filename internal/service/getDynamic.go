package service

import (
	"content/internal/db"
	"context"
	"github.com/huerni/gmitex/pkg/errno"

	"content/internal/svc"
	"content/pb"
)

type GetDynamicService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDynamicService(ctx context.Context, svcCtx *svc.ServiceContext) *GetDynamicService {
	return &GetDynamicService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *GetDynamicService) GetDynamic(in *pb.GetDynamicReq) (*pb.GetDynamicResp, error) {
	// todo: add your logic here and delete this line
	dynamicInfo := make([]*pb.DynamicInfo, 0)
	var dynamics []*db.Dynamic
	var err error
	if in.Id != nil {
		dynamics, err = db.QueryDynamicById(s.ctx, *in.Id)
	} else if in.UserId != nil {
		dynamics, err = db.QueryDynamicByUserId(s.ctx, *in.UserId)
	} else if in.PetId != nil {
		dynamics, err = db.QueryDynamicByPetId(s.ctx, *in.PetId)
	}
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	for _, info := range dynamics {
		dynamicInfo = append(dynamicInfo, &pb.DynamicInfo{
			Id:      uint64(info.ID),
			Title:   info.Title,
			Content: info.Content,
		})
	}
	return &pb.GetDynamicResp{DynamicInfo: dynamicInfo}, nil
}
