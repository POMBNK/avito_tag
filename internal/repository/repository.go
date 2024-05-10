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

type TagRepository interface {
	CreateTag(ctx context.Context, tag entity.Tag) (int, error)
	CreateTags(ctx context.Context, tag []entity.Tag) ([]int, error)
	FindExistingTags(ctx context.Context, tags []entity.Tag) (map[entity.Tag]struct{}, error)
}

type FeatureRepository interface {
	CreateFeature(ctx context.Context, feature entity.Feature) (int, error)
	FindByID(ctx context.Context, featureID int) (entity.Feature, error)
}
