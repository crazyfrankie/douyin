package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/app/user/biz"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/app/user/common/response"
)

type Handler struct {
	svc *service.UserService
}

func NewHandler(svc *service.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoute(r *gin.Engine) {
	userGroup := r.Group("api/user")
	{
		userGroup.POST("register", h.Register())
		userGroup.POST("login", h.Login())
		userGroup.GET("", h.UserInfo())
	}
}

func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerReq biz.RegisterReq
		if err := c.Bind(&registerReq); err != nil {
			return
		}

		token, err := h.svc.Register(c.Request.Context(), registerReq)
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", token)

		response.Success(c, nil)
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req biz.LoginReq
		if err := c.Bind(&req); err != nil {
			return
		}

		token, err := h.svc.Login(c.Request.Context(), req)
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", token)

		response.Success(c, nil)
	}
}

func (h *Handler) UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.svc.GetUserInfo(c.Request.Context(), claim["user_id"].(int64))
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
