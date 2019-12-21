package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vskut/twigo/pkg/common/entity"
)

func Test_UsersConstructor(t *testing.T) {
	t.Run("NewUsersRepositoryPostgreSQL", func(t *testing.T) {
		repo := NewUsersRepositoryPostgreSQL(&sql.DB{})

		assert.IsType(t, &UsersRepositoryPostgreSQL{}, repo, "should be TweetsRepositoryPostgreSQL")
	})

}

func Test_UsersSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewUsersRepositoryPostgreSQL(db)

	user := entity.User{
		Username: "username",
		Email:    "test@test.com",
		Password: "password",
	}

	t.Run("valid", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO users\\((.+)\\)").
			WithArgs(user.Username, user.Email, user.Password).WillReturnRows(rows)

		user, err := repo.Save(user)

		assert.NotEmpty(t, user.ID)
		assert.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		mock.ExpectQuery("^INSERT INTO users\\((.+)\\)").
			WillReturnError(fmt.Errorf("some sql/data error")).
			WithArgs(user.Username, user.Email, user.Password)

		user, err := repo.Save(user)

		assert.Empty(t, user.ID)
		assert.Error(t, err)
	})

}

func Test_UsersGetBy(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewUsersRepositoryPostgreSQL(db)

	user := entity.User{
		Username: "username",
		Email:    "test@test.com",
		Password: "password",
	}

	t.Run("email valid", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(1, user.Username, user.Email, user.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		user, err := repo.GetByEmail(user.Email)

		assert.NotEmpty(t, user.ID)
		assert.NoError(t, err)
	})

	t.Run("email invalid", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM users").
			WillReturnError(fmt.Errorf("some sql/data error"))
		user, err := repo.GetByEmail("invalid")

		assert.Empty(t, user.ID)
		assert.Error(t, err)
	})

	t.Run("username valid", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(1, user.Username, user.Email, user.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		user, err := repo.GetByUsername(user.Username)

		assert.NotEmpty(t, user.ID)
		assert.NoError(t, err)
	})

	t.Run("username invalid", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM users").
			WillReturnError(fmt.Errorf("some sql/data error"))
		user, err := repo.GetByUsername("invalid")

		assert.Empty(t, user.ID)
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

	repo := NewUsersRepositoryPostgreSQL(db)

	sourceUser := entity.User{
		ID:       1,
		Username: "source",
		Email:    "source@user.com",
		Password: "123456",
	}
	destinationUser := entity.User{
		ID:       2,
		Username: "destination",
		Email:    "destination@user.com",
		Password: "123456",
	}

	t.Run("valid", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		rowsSubscriptionsCount := sqlmock.NewRows([]string{"count(1)"}).
			AddRow(0)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(sourceUser.ID, destinationUser.ID).
			WillReturnRows(rowsSubscriptionsCount)

		rowInsertSubscription := sqlmock.NewRows([]string{"id"}).
			AddRow(1)
		mock.ExpectQuery("^INSERT INTO user_subscriptions\\((.+)\\)").
			WithArgs(sourceUser.ID, destinationUser.ID).WillReturnRows(rowInsertSubscription)

		err := repo.Subscribe(sourceUser, destinationUser.Username)
		assert.NoError(t, err)
	})

	t.Run("already subscribed", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		rowsSubscriptionsCount := sqlmock.NewRows([]string{"count(1)"}).
			AddRow(1)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(sourceUser.ID, destinationUser.ID).
			WillReturnRows(rowsSubscriptionsCount)

		err := repo.Subscribe(sourceUser, destinationUser.Username)
		assert.Error(t, err)
	})

	t.Run("select subscriptions count error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(sourceUser.ID, destinationUser.ID).
			WillReturnError(fmt.Errorf("some sql/data error"))

		err := repo.Subscribe(sourceUser, destinationUser.Username)
		assert.Error(t, err)
	})

	t.Run("subscription insert error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
			AddRow(destinationUser.ID, destinationUser.Username, destinationUser.Email, destinationUser.Password)
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		rowsSubscriptionsCount := sqlmock.NewRows([]string{"count(1)"}).
			AddRow(0)
		mock.ExpectQuery("^SELECT (.+) FROM user_subscriptions").
			WithArgs(sourceUser.ID, destinationUser.ID).
			WillReturnRows(rowsSubscriptionsCount)

		mock.ExpectQuery("^INSERT INTO user_subscriptions\\((.+)\\)").
			WithArgs(sourceUser.ID, destinationUser.ID).
			WillReturnError(fmt.Errorf("some sql/data error"))

		err := repo.Subscribe(sourceUser, destinationUser.Username)
		assert.Error(t, err)
	})

	t.Run("the same username", func(t *testing.T) {
		err := repo.Subscribe(sourceUser, sourceUser.Username)
		assert.Error(t, err)
	})

	t.Run("user with username doesn't exists", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "email", "password"})
		mock.ExpectQuery("^SELECT (.+) FROM users").WillReturnRows(rows)

		err := repo.Subscribe(sourceUser, "nonexistent-username")
		assert.Error(t, err)
	})
}
