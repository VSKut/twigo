package token

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vskut/twigo/pkg/common/entity"
)

// JwtTokenClaim ...
type JwtTokenClaim struct {
	User entity.User
	jwt.StandardClaims
}

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET"))

// CreateToken creates jwt token from user entity
func CreateToken(user entity.User) (string, error) {
	claims := JwtTokenClaim{
		User: entity.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "twigo-login",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ParseToken parses JWT token to user entity
func ParseToken(token string) (entity.User, error) {
	resultToken, err := jwt.ParseWithClaims(token, &JwtTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return entity.User{}, err
	}

	claims, ok := resultToken.Claims.(*JwtTokenClaim)
	if !ok || !resultToken.Valid {
		return entity.User{}, errors.New("invalid request: jwt-token is invalid")
	}

	return claims.User, nil
}

func JwtMiddleware(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := ParseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, "tokenInfo", tokenInfo), nil
}
