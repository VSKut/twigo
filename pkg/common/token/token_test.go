package token

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/vskut/twigo/pkg/common/entity"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func Test_Token(t *testing.T) {
	user := entity.User{
		ID:       1,
		Username: "TestUsername",
		Email:    "test@email.com",
	}

	resultToken, err := CreateToken(user)
	assert.NoError(t, err, "should be nil")

	t.Run("Create token", func(t *testing.T) {
		assert.NotEmpty(t, resultToken, "should be not empty")
		assert.NoError(t, err, "should be nil")
	})

}
func Test_ParseToken(t *testing.T) {
	user := entity.User{
		ID:       1,
		Username: "TestUsername",
		Email:    "test@email.com",
	}

	resultToken, err := CreateToken(user)
	assert.NoError(t, err, "should be nil")

	t.Run("Parse token", func(t *testing.T) {
		resultUser, err := ParseToken(resultToken)
		assert.Equal(t, user, resultUser, "should be equal as init user")
		assert.NoError(t, err, "should be nil")
	})

	t.Run("Token invalid sign", func(t *testing.T) {
		claims := JwtTokenClaim{
			User: user,
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

func Test_JwtMiddleware(t *testing.T) {
	user := entity.User{
		ID:       1,
		Username: "TestUsername",
		Email:    "test@email.com",
	}

	t.Run("valid", func(t *testing.T) {
		token, err := CreateToken(user)
		assert.NoError(t, err, "should be nil")

		ctx := context.Background()
		ctx = context.WithValue(ctx, ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+token),
		)

		resultCtx, err := JwtMiddleware(ctx)

		authUser, ok := resultCtx.Value(ValueTokenContextKey).(entity.User)

		assert.True(t, ok)
		assert.Equal(t, user, authUser)
		assert.NoError(t, err, "should be nil")
	})

	t.Run("Invalid token scheme", func(t *testing.T) {
		ctx := context.Background()
		_, err := JwtMiddleware(ctx)

		assert.Error(t, err, "should be nil")
	})

	t.Run("Invalid token", func(t *testing.T) {
		ctx := context.Background()
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer invalid-token-string"),
		)

		_, err := JwtMiddleware(ctx)

		assert.Error(t, err, "should be nil")
	})
}
