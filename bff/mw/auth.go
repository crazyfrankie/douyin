package mw

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	SecretKey = []byte("xT4vC0eM3tR2lP7kH8gD0kT0oN4oM6yX")
)

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

func (a *AuthBuild) Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := a.paths[r.URL.Path]; ok {
			next.ServeHTTP(w, r)
			return
		}

		tokenHeader := r.Header.Get("Authorization")
		token := extractToken(tokenHeader)

		_, err := parseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
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

func parseToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
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
