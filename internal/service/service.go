package service

import (
	"context"
	"github.com/POMBNK/avitotag/internal/entity"
)

// BannerService интерфейс для использования в сервисе
type BannerService interface {
	CreateBannerWithEntities(ctx context.Context, banner entity.Banner) (int, error)
	AllBanners(ctx context.Context, params entity.BannerSearchParams) ([]entity.Banner, error)
	UpdateBanner(ctx, bannerID int)
	DeleteBanner(ctx, bannerID int)
}
