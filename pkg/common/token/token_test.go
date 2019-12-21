package token

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"

	"github.com/vskut/twigo/pkg/common/entity"
)

func Test_Token(t *testing.T) {
	user := entity.User{
		ID:       1,
		Username: "TestUsername",
		Email:    "test@email.com",
	}

	resultToken, err := CreateToken(user)
	t.Run("Create token", func(t *testing.T) {
		assert.NotEmpty(t, resultToken, "should be not empty")
		assert.NoError(t, err, "should be nil")
	})

	t.Run("Parse token", func(t *testing.T) {
		resultUser, err := ParseToken(resultToken)
		assert.Equal(t, user, resultUser, "should be equal as init user")
		assert.NoError(t, err, "should be nil")
	})

	t.Run("Token invalid sign", func(t *testing.T) {
		claims := JwtTokenClaim{
			User: entity.User{
				ID:       1,
				Email:    "test@email.com",
				Username: "TestUser",
			},
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, _ := testToken.SignedString("invalid-sign-key")

		resultUser, err := ParseToken(token)
		assert.Equal(t, entity.User{}, resultUser, "should be empty")
		assert.Error(t, err, "should be error")
	})
}

func Test_ParseHeader(t *testing.T) {
	t.Run("Token is provided", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer 123123"),
		)
		token, err := ParseHeader(ctx)

		assert.NotEmpty(t, token, "shouldn't be empty")
		assert.NoError(t, err, "should be an error")
	})

	t.Run("Token is not provided in header", func(t *testing.T) {
		token, err := ParseHeader(context.Background())

		assert.Empty(t, token, "should be empty")
		assert.Error(t, err, "should be an error")
	})

	t.Run("Token is not provided Authorization in header", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("invalid", "invalid"),
		)
		token, err := ParseHeader(ctx)

		assert.Empty(t, token, "should be empty")
		assert.Error(t, err, "should be an error")
	})

	t.Run("Token is not provided Bearer in header", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "invalid"),
		)
		token, err := ParseHeader(ctx)

		assert.Empty(t, token, "should be empty")
		assert.Error(t, err, "should be an error")
	})

}

func Test_CheckAuth(t *testing.T) {
	t.Run("Token is provided", func(t *testing.T) {
		user := entity.User{
			ID:       1,
			Username: "TestUsername",
			Email:    "test@email.com",
		}

		resultToken, errResultToken := CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)
		token, user, err := CheckAuth(ctx)

		assert.NotEmpty(t, token, "shouldn't be empty")
		assert.NotEqual(t, entity.User{}, user, "should be not empty")
		assert.NoError(t, err, "shouldn't be an error")
	})

	t.Run("Token is not provided in header", func(t *testing.T) {
		token, user, err := CheckAuth(context.Background())

		assert.Empty(t, token, "should be empty")
		assert.Empty(t, user, "should be empty")
		assert.Error(t, err, "should be an error")
	})

	t.Run("Token is invalid", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer invalid-token"),
		)
		token, user, err := CheckAuth(ctx)

		assert.Empty(t, token, "shouldn't be empty")
		assert.Equal(t, entity.User{}, user, "shouldn't be empty")
		assert.Error(t, err, "should be nil")
	})
}
