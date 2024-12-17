package app

import (
	"context"
	"fmt"
	"homework8/internal/ads"
	"homework8/internal/user"
	"time"
)

type App struct {
	repoadapter ads.Repository
	useradapter user.Repository
}

// Конструктор приложения, который принимает репозитории для объявлений и пользователей
func NewApp(repo ads.Repository, user user.Repository) *App {
	return &App{repoadapter: repo,
		useradapter: user}
}

// Функция для создания пользователя. Принимает входные данные после валидации
func (a *App) CreateUser(ctx context.Context, nickname, email string) (*user.User, error) {
	user, err := a.useradapter.CreateUser(user.User{
		Id:       0,
		Nickname: nickname,
		Email:    email,
	})
	return user, err
}

// Функция для изменения пользователя. Принимает входные данные после валидации
func (a *App) UpdateUser(ctx context.Context, id int64, nickname, email string) (*user.User, error) {
	user, err := a.useradapter.UpdateUser(id, nickname, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Функция для получения пользователя по идентификатору
func (a *App) GetUser(ctx context.Context, id int64) (*user.User, error) {
	user, err := a.useradapter.GetUserbyId(id)
	return user, err
}

// Функция для получения всех объявлений пользователя. Ищет объявления во всех объявлениях по совпадению идентификатора пользователя
func (a *App) GetUserListAds(ctx context.Context, id int64) (*[]ads.Ad, error) {
	ads, err := a.repoadapter.GetUserAds(id)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

// Функция для создания объявления. Принимает входные данные после валидации
func (a *App) CreateAd(ctx context.Context, title string, text string, authorid int64) (*ads.Ad, error) {
	ad, err := a.repoadapter.CreateAd(ads.Ad{
		ID:        0,
		Title:     title,
		Text:      text,
		AuthorID:  authorid,
		Published: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return ad, err
}

// Функция для получения всего списка хранящихся объявлений
func (a *App) GetAds(ctx context.Context) (*[]ads.Ad, error) {
	ads, err := a.repoadapter.GetAllAds()
	if err != nil {
		return nil, fmt.Errorf("ErrorAdsNF")
	}
	return ads, nil
}

func (a *App) GetAdsByName(ctx context.Context, name string) (*[]ads.Ad, error) {
	ads, err := a.repoadapter.GetAdsByName(name)
	if err != nil {
		return nil, fmt.Errorf("ErrorAdsNF")
	}
	return ads, nil
}
func (a *App) GetAdByID(ctx context.Context, id int64) (*ads.Ad, error) {
	ad, err := a.repoadapter.GetAd(int64(id))
	if err != nil {
		return &ads.Ad{}, fmt.Errorf("ErrorAdNF")
	}
	return ad, nil
}

// Функция для изменения данных в объявлении. Проверяет что объявление изменяет автор.
func (a *App) UpdateAd(ctx context.Context, id int64, title string, text string, authorid int64) (*ads.Ad, error) {
	ad, err := a.repoadapter.UpdateInfoAd(id, authorid, title, text)
	if err != nil {
		return &ads.Ad{}, err
	}
	return ad, nil
}

// Функция для изменения статуса объявления
func (a *App) ChangeAdStatus(ctx context.Context, id int64, authorid int64, publised bool) (*ads.Ad, error) {
	ad, err := a.repoadapter.UpdateStatusAd(id, authorid, publised)
	if err != nil {
		return &ads.Ad{}, err
	}
	return ad, nil
}
