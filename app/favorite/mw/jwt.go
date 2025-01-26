package mw

import (
	"errors"
	
	"github.com/golang-jwt/jwt"

	"github.com/crazyfrankie/douyin/app/favorite/config"
)

var (
	SecretKey = []byte(config.GetConf().JWT.SecretKey)
)

func ParseToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.MapClaims); ok && tokenClaims.Valid {
			return *claims, nil
		}
	}

	return nil, errors.New("token is invalid")
}
