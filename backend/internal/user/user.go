package user

type User struct {
	Id       int64
	Nickname string `validate:"max:100"`
	Email    string `validate:"max:100"`
}
