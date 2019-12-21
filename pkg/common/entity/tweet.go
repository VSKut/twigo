package entity

// TweetsRepository ...
type TweetsRepository interface {
	Save(Tweet) (Tweet, error)
	ListAllByUser(User) ([]Tweet, error)
}

// Tweet ...
type Tweet struct {
	ID      uint64
	UserID  uint64
	Message string
}
