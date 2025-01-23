package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
)

type Handler struct {
	client publish.PublishServiceClient
}

func NewHandler(client publish.PublishServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	publishGroup := r.Group("api/publish")
	{
		publishGroup.POST("action", h.PublishAction())
		publishGroup.GET("list", h.PublishList())
	}
}

// PublishAction publish a video
//
// @router /douyin/publish/action/ [POST]
func (h *Handler) PublishAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PublishReq
		if err := c.Bind(&req); err != nil {
			response.Error(c, err)
			return
		}

		data, err := json.Marshal(req.Data)
		if err != nil {
			response.Error(c, err)
			return
		}

		_, err = h.client.PublishAction(c.Request.Context(), &publish.PublishActionRequest{
			Title: req.Title,
			Data:  data,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

// PublishList get the video list of user
//
// @router /douyin/publish/list/ [GET]
func (h *Handler) PublishList() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
