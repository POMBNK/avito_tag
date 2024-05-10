package feature

import (
	"context"
	"errors"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/pkg/apierror"
	"github.com/POMBNK/avitotag/internal/repository"
)

type Service struct {
	storage repository.FeatureRepository
}

func New(storage repository.FeatureRepository) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateFeature(ctx context.Context, feature entity.Feature) (int, error) {
	if feature.ID == 0 {
		//todo: logging, tracing
		return 0, apierror.ErrIvalidData
	}

	return s.storage.CreateFeature(ctx, feature)
}

func (s *Service) FindByID(ctx context.Context, featureID int) (entity.Feature, error) {
	return s.storage.FindByID(ctx, featureID)
}

func (s *Service) FindOrCreateFeature(ctx context.Context, feature entity.Feature) (int, error) {
	foundFeature, err := s.storage.FindByID(ctx, feature.ID)
	if err != nil {
		if !errors.Is(err, apierror.ErrNotFound) {
			return 0, err
		}
	}

	if foundFeature.ID != 0 {
		return foundFeature.ID, nil
	}

	return s.storage.CreateFeature(ctx, feature)
}
