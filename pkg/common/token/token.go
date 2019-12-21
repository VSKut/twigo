package token

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

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

// ParseHeader parses header from context.Context to string-token
func ParseHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("invalid request: jwt-token is not provided")
	}

	str, ok := md["authorization"]
	if !ok {
		return "", errors.New("invalid request: valid jwt-token required")
	}

	token := strings.Split(str[0], "Bearer ")
	if len(token) != 2 {
		return "", errors.New("invalid request: valid jwt-token required")
	}

	return token[1], nil
}

// CheckAuth checks authorization with jwt token
func CheckAuth(ctx context.Context) (string, entity.User, error) {
	jwtToken, err := ParseHeader(ctx)
	if err != nil {
		return "", entity.User{}, err
	}

	authUser, err := ParseToken(jwtToken)
	if err != nil {
		return "", entity.User{}, err
	}

	return jwtToken, authUser, nil
}
