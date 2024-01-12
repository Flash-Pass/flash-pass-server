package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

const (
	UserId     = "userId"
	OpenId     = "openId"
	SessionKey = "sessionKey"
	UnionId    = "unionId"
)

var (
	secret = "flash-pass"
)

type UserClaim struct {
	Id         int64  `json:"id"`
	OpenId     string `json:"open_id"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"union_id"`
	jwt.StandardClaims
}

func GenerateToken(ctx context.Context, claim *UserClaim) (string, error) {
	logger := ctxlog.Extract(ctx)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	if err != nil {
		logger.Error("generate token defeat", zap.Error(err), zap.Any("claim", claim))
		return "", err
	}

	return token, nil
}

func ParseToken(ctx context.Context, tokenString string) (*UserClaim, error) {
	logger := ctxlog.Extract(ctx)
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		logger.Error("token is invalid", zap.Any("token", tokenString))
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}

func ParseInfoWithToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	claim, err := ParseToken(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	result[UserId] = fmt.Sprint(claim.Id)
	result[OpenId] = claim.OpenId
	result[SessionKey] = claim.SessionKey
	result[UnionId] = claim.UnionId

	return result, nil
}
