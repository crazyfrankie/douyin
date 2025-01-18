//go:build wireinject

package ioc

import (
	"fmt"
	"github.com/crazyfrankie/douyin/app/favorite/biz/handler"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/crazyfrankie/douyin/app/favorite/config"
	"github.com/crazyfrankie/douyin/app/favorite/rpc/server"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
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

func InitWeb(favorite *handler.FavoriteHandler) *gin.Engine {
	engine := gin.Default()
	engine.Use(service.NewAuthBuilder().Auth())
	favorite.RegisterRoute(engine)

	return engine
}

func InitApp() *App {
	wire.Build(
		InitDB,

		dao.NewFavoriteDao,
		repository.NewFavoriteRepo,
		service.NewFavoriteService,
		handler.NewFavoriteHandler,

		server.NewFavoriteServer,

		InitWeb,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
