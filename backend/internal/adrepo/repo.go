package adrepo

import (
	"backend/internal/ads"
	"fmt"
	"sync"
	"time"
)

type RepositorySlice struct {
	mu  sync.Mutex
	ads []ads.Ad // slice
}

func New() *RepositorySlice {
	return &RepositorySlice{ads: make([]ads.Ad, 0)} // TODO: реализовать
}

func (r *RepositorySlice) CreateAd(ad ads.Ad) (*ads.Ad, error) {
	ad.ID = int64(len(r.ads)) // assuming ID is auto-incremented
	r.ads = append(r.ads, ad)
	return &ad, nil
}
func (r *RepositorySlice) GetAd(id int64) (*ads.Ad, error) {
	for _, ad := range r.ads {
		if ad.ID == id {
			return &ad, nil
		}
	}
	return &ads.Ad{}, fmt.Errorf("ad not found")
}
func (r *RepositorySlice) GetAllAds() (*[]ads.Ad, error) {
	return &r.ads, nil
}
func (r *RepositorySlice) UpdateInfoAd(id, userID int64, Title, Text string) (*ads.Ad, error) {
	for i, ad := range r.ads {
		if ad.ID == id {
			if userID != ad.AuthorID {
				return nil, fmt.Errorf("permission denied")
			}
			defer r.mu.Unlock()
			r.mu.Lock()
			r.ads[i].Title = Title
			r.ads[i].Text = Text
			r.ads[i].UpdatedAt = time.Now()
			return &r.ads[i], nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func (r *RepositorySlice) UpdateStatusAd(id int64, userID int64, status bool) (*ads.Ad, error) {
	for i, ad := range r.ads {
		if ad.ID == id {
			if userID != ad.AuthorID {
				return nil, fmt.Errorf("permission denided")
			}
			defer r.mu.Unlock()
			r.mu.Lock()
			r.ads[i].Published = status
			r.ads[i].UpdatedAt = time.Now()
			return &r.ads[i], nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func (r *RepositorySlice) GetUserAds(id int64) (*[]ads.Ad, error) {
	var userAds []ads.Ad
	for _, ad := range r.ads {
		if ad.AuthorID == id {
			userAds = append(userAds, ad)
		}
	}
	if len(userAds) == 0 {
		return nil, fmt.Errorf("user has no ads")
	}
	return &userAds, nil
}
func (r *RepositorySlice) GetAdsByName(name string) (*[]ads.Ad, error) {
	var adsByName []ads.Ad
	for _, ad := range r.ads {
		if ad.Title == name {
			adsByName = append(adsByName, ad)
		}
	}
	if len(adsByName) == 0 {
		return nil, fmt.Errorf("no ads found by this name")
	}
	return &adsByName, nil
}
