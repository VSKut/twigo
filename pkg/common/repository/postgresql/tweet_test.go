package postgresql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/vskut/twigo/pkg/common/entity"
)

func Test_TweetsConstructor(t *testing.T) {
	repo := NewTweetsRepositoryPostgreSQL(&sql.DB{})

	t.Run("NewTweetsRepositoryPostgreSQL", func(t *testing.T) {
		assert.IsType(t, &TweetsRepositoryPostgreSQL{}, repo, "should be TweetsRepositoryPostgreSQL")
	})
}

func Test_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewTweetsRepositoryPostgreSQL(db)

	t.Run("valid", func(t *testing.T) {
		tweet := entity.Tweet{UserID: 1, Message: "test message"}
		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO tweets\\((.+)\\)").
			WithArgs(tweet.Message, tweet.UserID).WillReturnRows(rows)

		tweet, err := repo.Save(tweet)

		assert.NotEmpty(t, tweet.ID)
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		tweet := entity.Tweet{UserID: 0, Message: ""}
		mock.ExpectQuery("^INSERT INTO tweets\\((.+)\\)").
			WillReturnError(fmt.Errorf("some sql/data error")).
			WithArgs(tweet.Message, tweet.UserID)

		tweet, err := repo.Save(tweet)

		assert.Empty(t, tweet.ID)
		assert.Error(t, err)
	})

}

func Test_TweetsListAllByUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewTweetsRepositoryPostgreSQL(db)

	t.Run("valid", func(t *testing.T) {
		user := entity.User{ID: 1}
		rows := sqlmock.NewRows([]string{"tweets.id", "tweets.message"}).
			AddRow(1, "tweet message 1").
			AddRow(2, "tweet message 2")
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").WillReturnRows(rows)

		tweets, err := repo.ListAllByUser(user)

		assert.Equal(t, len(tweets), 2)
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		user := entity.User{ID: 1}
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WillReturnError(fmt.Errorf("some sql/data error"))

		tweets, err := repo.ListAllByUser(user)

		assert.Nil(t, tweets)
		assert.Error(t, err)
	})

	t.Run("invalid tweet", func(t *testing.T) {
		user := entity.User{ID: 1}
		rows := sqlmock.NewRows([]string{"invalid"}).
			AddRow(nil)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").WillReturnRows(rows)

		tweets, err := repo.ListAllByUser(user)

		assert.Nil(t, tweets)
		assert.Error(t, err)
	})
}
