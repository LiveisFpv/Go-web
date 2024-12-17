package user

type Repository interface {
	GetUserbyId(id int64) (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(id int64, nickname string, email string) (*User, error)
}
