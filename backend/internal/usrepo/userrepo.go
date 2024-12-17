package usrepo

import (
	"fmt"
	"homework8/internal/user"
)

type RepoUser struct {
	users []user.User // slice
}

func New() *RepoUser {
	return &RepoUser{users: make([]user.User, 0)}
}
func (r *RepoUser) GetUserbyId(id int64) (*user.User, error) {
	for _, u := range r.users {
		if u.Id == id {
			return &u, nil
		}
	}
	return &user.User{}, fmt.Errorf("user not found %d", id)
}
func (r *RepoUser) CreateUser(user user.User) (*user.User, error) {
	user.Id = int64(len(r.users))
	r.users = append(r.users, user)
	return &user, nil
}
func (r *RepoUser) UpdateUser(id int64, nickname string, email string) (*user.User, error) {
	for i, u := range r.users {
		if u.Id == id {
			r.users[i].Nickname = nickname
			r.users[i].Email = email
			return &r.users[i], nil
		}
	}
	return &user.User{}, fmt.Errorf("user not found %d", id)
}
