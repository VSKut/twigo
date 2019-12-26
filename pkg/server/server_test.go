package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vskut/twigo/pkg/common/entity"
	"github.com/vskut/twigo/pkg/common/token"
	proto "github.com/vskut/twigo/pkg/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_ServerConstructor(t *testing.T) {
	srv := NewServer(&sql.DB{})
	assert.IsType(t, &Server{}, srv, "should be TweetsRepositoryPosNewServertgreSQL")
}

func Test_Run(t *testing.T) {
	srv := NewServer(&sql.DB{})

	t.Run("invalid port", func(t *testing.T) {
		err := srv.Run("invalid")

		assert.Error(t, err)
	})
}

func Test_AuthFuncOverride(t *testing.T) {
	srv := NewServer(&sql.DB{})

	t.Run("need token", func(t *testing.T) {
		user := entity.User{
			ID:       1,
			Username: "username",
			Email:    "test@test.com",
			Password: "123456",
		}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)
		_, err := srv.AuthFuncOverride(ctx, "/tweet.TweetService/ListTweet")

		assert.NoError(t, err)
	})

	t.Run("no need token", func(t *testing.T) {
		ctx := context.Background()
		_, err := srv.AuthFuncOverride(ctx, "/auth.AuthService/Login")

		assert.NoError(t, err)
	})
}

func Test_Login(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	srv := NewServer(db)

	ctx := context.Background()

	t.Run("valid", func(t *testing.T) {
		request := &proto.LoginRequest{
			Email:    "test@test.com",
			Password: "123456",
		}

		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(1, "test", request.Email, "$2a$10$g6isZT.FNtWbQfIvlw5j.OfkcSNLYBqDngvwcvGrQQmownL4JKUBO")
		mock.ExpectQuery("^SELECT (.+) FROM users").
			WithArgs(request.Email).
			WillReturnRows(rows)

		_, err := srv.Login(ctx, request)

		assert.NoError(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.LoginRequest{}

		_, err = srv.Login(ctx, request)

		assert.Error(t, err)
	})

	t.Run("nonexistent email provided", func(t *testing.T) {
		request := &proto.LoginRequest{
			Email:    "nonexistent@test.com",
			Password: "123456",
		}

		mock.ExpectQuery("^SELECT (.+) FROM users").
			WithArgs(request.Email).
			WillReturnError(fmt.Errorf("no records found error"))

		resp, err := srv.Login(ctx, request)

		assert.Empty(t, resp.Token)
		assert.Error(t, err)
	})

	t.Run("invalid password provided", func(t *testing.T) {
		request := &proto.LoginRequest{
			Email:    "test@test.com",
			Password: "123456",
		}

		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(1, "test", request.Email, "hash-of-invalid-password")
		mock.ExpectQuery("^SELECT (.+) FROM users").
			WithArgs(request.Email).
			WillReturnRows(rows)

		_, err := srv.Login(ctx, request)

		assert.Error(t, err)
	})
}

func Test_Register(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	srv := NewServer(db)

	ctx := context.Background()

	t.Run("valid", func(t *testing.T) {
		request := &proto.RegisterRequest{
			Email:    "test@test.com",
			Password: "123456",
			Username: "username",
		}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO users\\((.+)\\)").WillReturnRows(rows)

		_, err = srv.Register(ctx, request)

		assert.NoError(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.RegisterRequest{}

		_, err = srv.Register(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid save repo response", func(t *testing.T) {
		request := &proto.RegisterRequest{
			Email:    "test@test.com",
			Password: "123456",
			Username: "username",
		}

		mock.ExpectQuery("^INSERT INTO users\\((.+)\\)").
			WillReturnError(fmt.Errorf("some sql error"))

		_, err = srv.Register(ctx, request)

		assert.Error(t, err)
	})
}

func Test_Subscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	srv := NewServer(db)

	user := entity.User{
		ID:       1,
		Username: "username",
		Email:    "test@test.com",
		Password: "123456",
	}

	t.Run("valid", func(t *testing.T) {
		destinationUser := entity.User{
			ID:       2,
			Username: "destination",
			Email:    "destination@user.com",
			Password: "123456",
		}

		request := &proto.SubscribeRequest{
			Username: destinationUser.Username,
		}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		rowsSubscriptionsCount := sqlmock.NewRows([]string{"count(1)"}).
			AddRow(0)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(user.ID, destinationUser.ID).
			WillReturnRows(rowsSubscriptionsCount)

		rowInsertSubscription := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO user_subscriptions\\((.+)\\)").
			WithArgs(user.ID, destinationUser.ID).WillReturnRows(rowInsertSubscription)

		_, err = srv.Subscribe(ctx, request)

		assert.NoError(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.SubscribeRequest{}

		ctx := context.Background()

		_, err = srv.Subscribe(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid auth validation", func(t *testing.T) {
		request := &proto.SubscribeRequest{
			Username: "existing-username",
		}

		ctx := context.Background()

		_, err = srv.Subscribe(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		destinationUser := entity.User{
			ID:       2,
			Username: "destination",
			Email:    "destination@user.com",
			Password: "123456",
		}

		request := &proto.SubscribeRequest{
			Username: "username",
		}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		rowsSubscriptionsCount := sqlmock.NewRows([]string{"count(1)"}).
			AddRow(0)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(user.ID, destinationUser.ID).
			WillReturnRows(rowsSubscriptionsCount)

		mock.ExpectQuery("^INSERT INTO user_subscriptions\\((.+)\\)").
			WithArgs(user.ID, destinationUser.ID).
			WillReturnError(fmt.Errorf("some sql error"))

		_, err = srv.Subscribe(ctx, request)

		assert.Error(t, err)
	})
}

func Test_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	srv := NewServer(db)

	ctx := context.Background()

	user := entity.User{
		ID:       1,
		Username: "username",
		Email:    "test@test.com",
		Password: "123456",
	}

	t.Run("valid", func(t *testing.T) {
		request := &proto.CreateTweetRequest{
			Message: "message",
		}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO tweets\\((.+)\\)").
			WithArgs(request.Message, 1).WillReturnRows(rows)

		_, err = srv.CreateTweet(ctx, request)

		assert.NoError(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.CreateTweetRequest{}

		_, err = srv.CreateTweet(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid auth validation", func(t *testing.T) {
		request := &proto.CreateTweetRequest{
			Message: "message",
		}

		ctx := context.Background()

		_, err = srv.CreateTweet(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.CreateTweetRequest{
			Message: "message",
		}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		mock.ExpectQuery("^INSERT INTO tweets\\((.+)\\)").
			WithArgs(request.Message, 1).
			WillReturnError(fmt.Errorf("some sql error"))

		_, err = srv.CreateTweet(ctx, request)

		assert.Error(t, err)
	})
}

func Test_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	srv := NewServer(db)

	//ctx := context.Background()

	user := entity.User{
		ID:       1,
		Username: "username",
		Email:    "test@test.com",
		Password: "123456",
	}

	t.Run("invalid auth validation", func(t *testing.T) {
		request := &proto.ListTweetRequest{}

		ctx := context.Background()

		_, err = srv.ListTweet(ctx, request)

		assert.Error(t, err)
	})

	t.Run("invalid request validation", func(t *testing.T) {
		request := &proto.ListTweetRequest{}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WillReturnError(fmt.Errorf("some sql error"))

		_, err = srv.ListTweet(ctx, request)

		assert.Error(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		request := &proto.ListTweetRequest{}

		resultToken, errResultToken := token.CreateToken(user)

		assert.NotEmpty(t, resultToken, "shouldn't be empty")
		assert.NoError(t, errResultToken, "shouldn't be an error")

		ctx := context.Background()
		ctx = context.WithValue(ctx, token.ValueTokenContextKey, user)
		ctx = metadata.NewIncomingContext(
			ctx,
			metadata.Pairs("authorization", "Bearer "+resultToken),
		)

		rows := sqlmock.NewRows([]string{"id", "message"}).
			AddRow(1, "message")
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WillReturnRows(rows)

		_, err = srv.ListTweet(ctx, request)

		assert.NoError(t, err)
	})

}
