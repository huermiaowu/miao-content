package db

import (
	"github.com/huermiaowu/miao-content/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Mysql.DSN))
	if err != nil {
		return err
	}

	// 根据结构自动建表
	err = DB.AutoMigrate(&Dynamic{})
	if err != nil {
		return err
	}
	return nil
}
