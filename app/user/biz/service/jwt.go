package service

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/app/user/config"
)

func GenerateToken(uid int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   uid,
		"expire_at": time.Now().Add(time.Hour * 24).Unix(),
		"issuer":    "github.com/crazyfrankie",
		"issue_at":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConf().JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
