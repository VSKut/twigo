package postgresql

import (
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/labstack/gommon/log"
	"github.com/vskut/twigo/pkg/common/entity"
)

// UsersRepositoryPostgreSQL ...
type UsersRepositoryPostgreSQL struct {
	db *sql.DB
}

// NewUsersRepositoryPostgreSQL constructs the UsersRepositoryPostgreSQL struct
func NewUsersRepositoryPostgreSQL(db *sql.DB) *UsersRepositoryPostgreSQL {
	return &UsersRepositoryPostgreSQL{
		db: db,
	}
}

// Save implements logic of saving user to db
func (r *UsersRepositoryPostgreSQL) Save(user entity.User) (entity.User, error) {
	query := "INSERT INTO users(username,email,password) VALUES($1,$2,$3) returning id;"
	err := r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Print(err)
		return entity.User{}, status.Error(codes.InvalidArgument, "invalid SaveRequest: request error")
	}

	return user, nil
}

// Subscribe implements logic of subscribing user to another user
func (r *UsersRepositoryPostgreSQL) Subscribe(userSource entity.User, username string) error {
	if userSource.Username == username {
		return status.Errorf(codes.InvalidArgument, "invalid SubscribeRequest.Username: user with username %q can't subscribe to himself", userSource.Username)
	}

	var user entity.User
	querySelectUserByUsername := "SELECT id, username, email, password FROM users WHERE username = $1;"
	err := r.db.QueryRow(querySelectUserByUsername, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil || user.ID == 0 {
		log.Print(err)
		return status.Errorf(codes.InvalidArgument, "invalid SubscribeRequest.Username: user with username %q doesn't exists", username)
	}

	var userSubscriptionExists uint
	querySelectSubscriptionsCount := "SELECT count(1) FROM user_subscriptions WHERE user_id = $1 AND destination_user_id = $2;"
	err = r.db.QueryRow(querySelectSubscriptionsCount, userSource.ID, user.ID).Scan(&userSubscriptionExists)
	if err != nil {
		log.Print(err)
		return status.Error(codes.InvalidArgument, "invalid SubscribeRequest.Username: request error")
	}

	if userSubscriptionExists > 0 {
		return status.Errorf(codes.InvalidArgument, "invalid SubscribeRequest.Username: user with username %q already subscribed to user %q", userSource.Username, username)
	}

	queryInsertSubscription := "INSERT INTO user_subscriptions(user_id,destination_user_id) VALUES($1,$2);"
	_, err = r.db.Query(queryInsertSubscription, userSource.ID, user.ID)
	if err != nil {
		log.Print(err)
		return status.Error(codes.InvalidArgument, "invalid SubscribeRequest.Username: request error")
	}

	return nil
}

// GetByEmail implements logic of retrieving user by email
func (r *UsersRepositoryPostgreSQL) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, email, password FROM users WHERE email = $1;"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Print(err)
		return entity.User{}, status.Errorf(codes.InvalidArgument, "invalid GetByEmail.Email: user with email %q doesn't exists", email)
	}

	return user, nil
}

// GetByUsername implements logic of retrieving user by username
func (r *UsersRepositoryPostgreSQL) GetByUsername(username string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, email, password FROM users WHERE username = $1;"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Print(err)
		return entity.User{}, status.Errorf(codes.InvalidArgument, "invalid GetByUsername.Username: user with username %q doesn't exists", username)
	}

	return user, nil
}
