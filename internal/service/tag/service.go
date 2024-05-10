package tag

import (
	"context"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/repository"
)

type Service struct {
	storage repository.TagRepository
}

func (s *Service) FindExistingTags(ctx context.Context, tags []entity.Tag) (map[entity.Tag]struct{}, error) {
	return s.storage.FindExistingTags(ctx, tags)
}

func New(storage repository.TagRepository) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateTag(ctx context.Context, tag entity.Tag) (int, error) {
	return s.storage.CreateTag(ctx, tag)
}

func (s *Service) CreateTags(ctx context.Context, tags []entity.Tag) ([]int, error) {
	return s.storage.CreateTags(ctx, tags)
}
