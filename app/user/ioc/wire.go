//go:build wireinject

package ioc

import (
	"fmt"
	"github.com/crazyfrankie/douyin/app/user/biz/handler"
	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/app/user/config"
	"github.com/crazyfrankie/douyin/app/user/rpc/server"
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

func InitWeb(md []gin.HandlerFunc, user *handler.Handler) *gin.Engine {
	engine := gin.Default()
	engine.Use(md...)
	user.RegisterRoute(engine)

	return engine
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		service.NewAuthBuilder().
			IgnorePath("/api/user/login").
			IgnorePath("/api/user/signup").
			Auth(),
	}
}

func InitApp() *App {
	wire.Build(
		InitDB,

		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewHandler,

		server.NewUserServer,

		InitMiddlewares,
		InitWeb,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
