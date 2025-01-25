package handler

import (
	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/rpc_gen/relation"
)

type Handler struct {
	client relation.RelationServiceClient
}

func NewHandler(client relation.RelationServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	relationGroup := r.Group("api/relation")
	{
		relationGroup.POST("action", h.RelationAction())
		relationGroup.GET("follow/list", h.RelationFollowList())
		relationGroup.GET("follower/list", h.RelationFollowerList())
		relationGroup.GET("friend/list", h.RelationFriendList())
	}
}

func (h *Handler) RelationAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RelationActionReq
		if err := c.Bind(&req); err != nil {
			return
		}

		_, err := h.client.RelationAction(c.Request.Context(), &relation.RelationActionRequest{
			UserId:     req.UserId,
			ToUserId:   req.ToUserId,
			ActionType: req.ActionType,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *Handler) RelationFollowList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.client.RelationFollowList(c.Request.Context(), &relation.RelationFollowListRequest{
			UserId: claim["user_id"].(int64),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *Handler) RelationFollowerList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.client.RelationFollowerList(c.Request.Context(), &relation.RelationFollowerListRequest{
			UserId: claim["user_id"].(int64),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *Handler) RelationFriendList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.client.RelationFriendList(c.Request.Context(), &relation.RelationFriendListRequest{
			UserId: claim["user_id"].(int64),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
