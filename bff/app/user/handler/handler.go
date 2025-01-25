package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/bff/common/response"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type Handler struct {
	client user.UserServiceClient
}

func NewHandler(client user.UserServiceClient) *Handler {
	return &Handler{client: client}
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
		var req RegisterReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.client.Register(c.Request.Context(), &user.RegisterRequest{
			Name:     req.Name,
			Password: req.Password,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", resp.GetToken())

		response.Success(c, nil)
	}
}

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.client.Login(c.Request.Context(), &user.LoginRequest{
			Name:     req.Name,
			Password: req.Password,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", resp.GetToken())

		response.Success(c, nil)
	}
}

func (h *Handler) UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(jwt.MapClaims)

		resp, err := h.client.GetUserInfo(c.Request.Context(), &user.GetUserInfoRequest{
			UserIdToQuery: claim["user_id"].(int64),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
