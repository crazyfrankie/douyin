package handler

import (
	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	client feed.FeedServiceClient
}

func NewHandler(client feed.FeedServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	feedGroup := r.Group("api/feed")
	{
		feedGroup.GET("", h.Feed())
	}
}

// Feed get a list of recommended videos
//
// @router /api/feed/ [GET]
func (h *Handler) Feed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var feedReq FeedReq
		if err := c.Bind(&feedReq); err != nil {
			return
		}

		resp, err := h.client.Feed(c.Request.Context(), &feed.FeedRequest{
			LatestTime: feedReq.LatestTime,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
