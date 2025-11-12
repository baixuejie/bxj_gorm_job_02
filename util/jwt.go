package utils

import (
	"bxj_gorm_job_02/config"
	"bxj_gorm_job_02/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	Username string
	NickName string
}

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

// 生成 JWT Token
func GenerateToken(u *model.User) (string, time.Time, error) {
	// 设置过期时间（例如 2 小时后过期）
	expirationTime := time.Now().Add(2 * time.Hour)

	// 构建 Claims
	claims := CustomClaims{
		BaseClaims: BaseClaims{
			UUID:     uuid.New(),
			ID:       u.Id,
			Username: u.Username,
			NickName: u.Username,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),     // 生效时间（立即生效）
			Issuer:    "your-app-name",                    // 签发者（可选）
			Subject:   "user-token",                       // 主题（可选）
		},
	}

	// 创建 Token（使用 HS256 签名算法）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成最终 Token 字符串
	tokenString, err := token.SignedString(config.CONFIG.Jwt.SigningKey)
	if err != nil {
		return "", expirationTime, fmt.Errorf("生成 Token 失败: %v", err)
	}
	return tokenString, expirationTime, nil
}

func CheckToken(token string) (isLogin bool, err error) {
	redisToken, err := config.REDIS.Get(context.Background(), token).Result()
	if err != nil {
		return false, err
	}
	if redisToken == "" {
		return false, errors.New("token不存在")
	}
	return true, nil
}
