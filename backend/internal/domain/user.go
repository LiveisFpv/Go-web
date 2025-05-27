package domain

type User struct {
	Id         int64
	Email      string
	Login      string
	Student_id *int
	Password   string
	Role       *string
}
