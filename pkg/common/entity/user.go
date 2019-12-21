package entity

// UsersRepository ...
type UsersRepository interface {
	Save(User) (User, error)
	Subscribe(User, string) error
	GetByEmail(string) (User, error)
	GetByUsername(string) (User, error)
}

// User ...
type User struct {
	ID       uint64
	Email    string
	Username string
	Password string `json:"password,omitempty"`
}
