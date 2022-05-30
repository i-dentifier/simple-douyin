package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-douyin/model"
)

var DB *gorm.DB

type Mysql struct {
	Root      string
	Pwd       string
	Host      string
	Port      string
	Database  string
	Charset   string
	ParseTime string
	Loc       string
}

func DBInit() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%s",
		// 数据库用户名
		"douyin",
		// 数据库密码
		"douyin100@",
		"180.76.52.150",
		// 库名
		"simple_douyin",
		"utf8mb4",
		"True",
		"Local",
		"10s",
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect mysql")
	}
	if err = DB.AutoMigrate(&model.User{}); err != nil {
		panic("failed to auto migrate database users")
	}

	if err = DB.AutoMigrate(&model.UserAuth{}); err != nil {
		panic("failed to auto migrate database user_auths")
	}

	if err = DB.AutoMigrate(&model.Video{}); err != nil {
		panic("failed to auto migrate database videos")
	}

	if err = DB.AutoMigrate(&model.Favorite{}); err != nil {
		panic("failed to auto migrate database favorite")
	}
	return nil
}
