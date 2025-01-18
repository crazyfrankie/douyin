// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"fmt"
	"github.com/crazyfrankie/douyin/app/feed/biz/handler"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/crazyfrankie/douyin/app/feed/config"
	"github.com/crazyfrankie/douyin/app/feed/rpc/server"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

// Injectors from wire.go:

func InitApp() *App {
	db := InitDB()
	feedDao := dao.NewFeedDao(db)
	feedRepo := repository.NewFeedRepo(feedDao)
	feedService := service.NewFeedService(feedRepo)
	feedHandler := handler.NewFeedHandler(feedService)
	engine := InitWeb(feedHandler)
	videoServer := server.NewVideoServer(feedService)
	app := &App{
		HTTPServer: engine,
		RPCServer:  videoServer,
	}
	return app
}

// wire.go:

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(config.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	return db
}

func InitWeb(feed *handler.FeedHandler) *gin.Engine {
	engine := gin.Default()
	engine.Use(service.NewAuthBuilder().Auth())
	feed.RegisterRoute(engine)

	return engine
}
