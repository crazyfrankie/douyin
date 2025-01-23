package handler

import (
	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	client favorite.FavoriteServiceClient
}

func NewHandler(client favorite.FavoriteServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	favoriteGroup := r.Group("api/favorite")
	{
		favoriteGroup.POST("action", h.FavoriteAction())
		favoriteGroup.GET("list", h.FavoriteList())
	}
}

func (h *Handler) FavoriteAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req FavoriteActionReq
		if err := c.Bind(&req); err != nil {
			return
		}

		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		_, err := h.client.FavoriteAction(c.Request.Context(), &favorite.FavoriteActionRequest{
			UserId:  claim["user_id"].(int64),
			VideoId: req.VideoId,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *Handler) FavoriteList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.client.FavoriteList(c.Request.Context(), &favorite.FavoriteListRequest{
			UserId: claim["user_id"].(int64),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
