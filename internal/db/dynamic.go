package db

import (
	"context"
	"gorm.io/gorm"
)

var dynamicTableName = "dynamic"

type Dynamic struct {
	gorm.Model
	Title   string `gorm:"column:title;type:varchar(254);not null"`
	Content string `gorm:"column:content;type:text;not null"`
	UserId  uint64 `gorm:"column:user_id;type:bigint;not null"`
	PetId   uint64 `gorm:"column:pet_id;type:bigint;"`
}

func (d *Dynamic) TableName() string {
	return dynamicTableName
}

func CreateDynamic(ctx context.Context, dynamic *Dynamic) error {
	return DB.WithContext(ctx).Create(dynamic).Error
}

func UpdateDynamic(ctx context.Context, dynamic *Dynamic) error {
	return DB.WithContext(ctx).Updates(dynamic).Error
}

func QueryDynamicById(ctx context.Context, ID uint64) ([]*Dynamic, error) {
	res := make([]*Dynamic, 0)
	if err := DB.WithContext(ctx).Where("id = ?", ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryDynamicByUserId(ctx context.Context, userId uint64) ([]*Dynamic, error) {
	res := make([]*Dynamic, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryDynamicByPetId(ctx context.Context, petId uint64) ([]*Dynamic, error) {
	res := make([]*Dynamic, 0)
	if err := DB.WithContext(ctx).Where("pet_id = ?", petId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
