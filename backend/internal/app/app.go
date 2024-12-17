package app

import (
	"backend/internal/repository"
)

type App struct {
	repo repository.Repository
}

// Конструктор приложения, который принимает репозитории для объявлений и пользователей
func NewApp(repo repository.Repository) *App {
	return &App{repo: repo}
}

// Функция для создания пользователя. Принимает входные данные после валидации
// func (a *App) CreateUser(ctx context.Context, nickname, email string) (*user.User, error) {
// 	user, err := a.useradapter.CreateUser(user.User{
// 		Id:       0,
// 		Nickname: nickname,
// 		Email:    email,
// 	})
// 	return user, err
// }
