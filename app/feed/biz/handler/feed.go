package handler

import (
	"github.com/gin-gonic/gin"
	
	"github.com/crazyfrankie/douyin/app/feed/biz"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
)

type FeedHandler struct {
	svc *service.FeedService
}

func NewFeedHandler(svc *service.FeedService) *FeedHandler {
	return &FeedHandler{svc: svc}
}

func (h *FeedHandler) RegisterRoute(r *gin.Engine) {
	feedGroup := r.Group("api/feed")
	{
		feedGroup.GET("", h.Feed())
	}
}

func (h *FeedHandler) Feed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var feedReq biz.FeedReq
		if err := c.Bind(&feedReq); err != nil {
			return
		}

	}
}
