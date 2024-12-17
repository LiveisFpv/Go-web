package ads

type Repository interface {
	GetUserAds(id int64) (*[]Ad, error)
	CreateAd(ad Ad) (*Ad, error)
	GetAd(id int64) (*Ad, error)
	GetAllAds() (*[]Ad, error)
	UpdateStatusAd(id int64, userID int64, status bool) (*Ad, error)
	UpdateInfoAd(id, userID int64, Title, Text string) (*Ad, error)
	GetAdsByName(name string) (*[]Ad, error)
}
