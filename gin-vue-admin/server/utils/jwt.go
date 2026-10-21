package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/radish-go/gin-vue-admin/server/global"
)

// CustomClaims 自定义载荷
type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 签发 JWT
func GenerateToken(id uint, username string) (string, error) {
	cfg := global.GVA_CONFIG.JWT
	expires, err := time.ParseDuration(cfg.ExpiresTime)
	if err != nil || expires <= 0 {
		expires = 24 * time.Hour
	}
	claims := CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.SigningKey))
}

// ParseToken 解析并校验 JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(global.GVA_CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
