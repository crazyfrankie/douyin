package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/app/favorite/biz"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/crazyfrankie/douyin/app/favorite/common/response"
)

type FavoriteHandler struct {
	svc *service.FavoriteService
}

func NewFavoriteHandler(svc *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{svc: svc}
}

func (h *FavoriteHandler) RegisterRoute(r *gin.Engine) {
	favoriteGroup := r.Group("api/favorite")
	{
		favoriteGroup.POST("action", h.FavoriteAction())
		favoriteGroup.GET("list", h.FavoriteList())
	}
}

func (h *FavoriteHandler) FavoriteAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var favoriteReq biz.FavoriteActionReq
		if err := c.Bind(&favoriteReq); err != nil {
			return
		}

		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		err := h.svc.FavoriteAction(c.Request.Context(), favoriteReq, claim["user_id"].(int64))
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *FavoriteHandler) FavoriteList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.svc.FavoriteList(c.Request.Context(), claim["user_id"].(int64))
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
