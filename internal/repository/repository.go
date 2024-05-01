package repository

import (
	"context"
	"github.com/POMBNK/avitotag/internal/entity"
)

type BannerRepository interface {
	CreateBanner(ctx context.Context, banner entity.Banner) (int, error)
	AllBanners()
	UpdateBanner()
	DeleteBanner()
}
