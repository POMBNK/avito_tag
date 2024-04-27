package service

// UserService интерфейс для использования в сервисе
type UserService interface {
	CreateBanner()
	AllBanners()
	UpdateBanner()
	DeleteBanner()
}
