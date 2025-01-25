package handler

import (
	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/comment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	client comment.CommentServiceClient
}

func NewHandler(client comment.CommentServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	commentGroup := r.Group("api/comment")
	{
		commentGroup.POST("action", h.CommentAction())
		commentGroup.GET("list", h.CommentList())
	}
}

func (h *Handler) CommentAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CommentReq
		if err := c.Bind(&req); err != nil {
			return
		}
		resp, err := h.client.CommentAction(c.Request.Context(), &comment.CommentActionRequest{
			VideoId:     req.VideoID,
			ActionType:  req.ActionType,
			CommentText: req.CommentText,
			CommentId:   req.CommentID,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *Handler) CommentList() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
