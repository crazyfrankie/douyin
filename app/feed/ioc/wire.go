//go:build wireinject

package ioc

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/crazyfrankie/douyin/app/feed/biz/handler"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/crazyfrankie/douyin/app/feed/config"
	"github.com/crazyfrankie/douyin/app/feed/rpc/server"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(config.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"))

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

func InitApp() *App {
	wire.Build(
		InitDB,

		dao.NewFeedDao,
		repository.NewFeedRepo,
		service.NewFeedService,
		handler.NewFeedHandler,
		server.NewVideoServer,

		InitWeb,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
