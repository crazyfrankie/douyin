package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/app/favorite/config"
)

func parseToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetConf().JWT.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, errors.New("token is invalid")
}

type AuthBuild struct {
	paths map[string]struct{}
}

func NewAuthBuilder() *AuthBuild {
	return &AuthBuild{paths: make(map[string]struct{})}
}

func (a *AuthBuild) IgnorePath(path string) *AuthBuild {
	a.paths[path] = struct{}{}
	return a
}

func (a *AuthBuild) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := a.paths[c.Request.URL.Path]; ok {
			c.Next()
			return
		}

		tokenHeader := c.GetHeader("Authorization")
		token := extractToken(tokenHeader)

		claims, err := parseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}

func extractToken(token string) string {
	if token == "" {
		return ""
	}

	strs := strings.Split(token, " ")
	if strs[0] == "Bearer" {
		return strs[1]
	}

	return ""
}
