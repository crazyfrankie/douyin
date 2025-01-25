package handler

import (
	"github.com/gin-gonic/gin"
	
	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/message"
)

type Handler struct {
	client message.MessageServiceClient
}

func NewHandler(client message.MessageServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	messageGroup := r.Group("api/message")
	{
		messageGroup.POST("action", h.MessageAction())
		messageGroup.GET("chat", h.MessageChat())
	}
}

func (h *Handler) MessageAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req MessageActionReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.client.MessageAction(c.Request.Context(), &message.MessageActionRequest{
			UserId:     req.UserId,
			ToUserId:   req.ToUserId,
			Content:    req.Content,
			ActionType: req.ActionType,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *Handler) MessageChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req MessageChatReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.client.MessageChat(c.Request.Context(), &message.MessageChatRequest{
			UserId:     req.UserId,
			ToUserId:   req.ToUserId,
			PreMsgTime: req.PreMsgTime,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
