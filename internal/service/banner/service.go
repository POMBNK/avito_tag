package banner

import (
	"context"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/repository"
)

type Service struct {
	storage repository.BannerRepository
}

func (s *Service) AllBanners(ctx context.Context, bannerID int) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateBanner(ctx, bannerID int) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteBanner(ctx, bannerID int) {
	//TODO implement me
	panic("implement me")
}

func New(storage repository.BannerRepository) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateBanner(ctx context.Context, banner entity.Banner) (int, error) {
	bannerID, err := s.storage.CreateBanner(ctx, banner)
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}
