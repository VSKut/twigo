package postgresql

import (
	"database/sql"
	"github.com/vskut/twigo/pkg/common/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// TweetsRepositoryPostgreSQL ...
type TweetsRepositoryPostgreSQL struct {
	db *sql.DB
}

// NewTweetsRepositoryPostgreSQL constructs the TweetsRepositoryPostgreSQL struct
func NewTweetsRepositoryPostgreSQL(db *sql.DB) *TweetsRepositoryPostgreSQL {
	return &TweetsRepositoryPostgreSQL{
		db: db,
	}
}

// Save implements logic of saving tweet to db
func (r *TweetsRepositoryPostgreSQL) Save(tweet entity.Tweet) (entity.Tweet, error) {
	query := "INSERT INTO tweets(message,user_id) VALUES($1,$2) returning id;"
	err := r.db.QueryRow(query, tweet.Message, tweet.UserID).Scan(&tweet.ID)
	if err != nil {
		log.Print(err)
		return entity.Tweet{}, status.Error(codes.InvalidArgument, "invalid SaveRequest: request error")
	}

	return tweet, nil
}

// ListAllByUser implements logic of listing all tweets by user
func (r *TweetsRepositoryPostgreSQL) ListAllByUser(user entity.User) ([]entity.Tweet, error) {
	var result []entity.Tweet

	query := "SELECT tweets.id, tweets.message FROM user_subscriptions INNER JOIN tweets ON (tweets.user_id = user_subscriptions.destination_user_id) WHERE user_subscriptions.user_id = $1;"
	rows, err := r.db.Query(query, user.ID)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.InvalidArgument, "invalid ListAllByUser: request error")
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tweet entity.Tweet
		err = rows.Scan(&tweet.ID, &tweet.Message)
		if err != nil {
			log.Print(err)
			return nil, status.Error(codes.InvalidArgument, "invalid ListAllByUser: request error")
		}
		result = append(result, tweet)
	}

	return result, nil
}
