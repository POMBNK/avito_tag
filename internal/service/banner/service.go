package banner

import (
	"context"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/repository"
	"github.com/POMBNK/avitotag/internal/repository/tx"
	"go.openly.dev/pointy"
)

type Tagger interface {
	CreateTags(ctx context.Context, tags []entity.Tag) ([]int, error)
	FindExistingTags(ctx context.Context, tags []entity.Tag) (map[entity.Tag]struct{}, error)
}

type Featter interface {
	FindOrCreateFeature(ctx context.Context, feature entity.Feature) (int, error)
}

type Service struct {
	tagService     Tagger
	featureService Featter
	storage        repository.BannerRepository
	tx             tx.Tx
}

func New(
	storage repository.BannerRepository,
	tagService Tagger,
	featureService Featter,
	tx tx.Tx,

) *Service {

	return &Service{
		storage:        storage,
		tagService:     tagService,
		featureService: featureService,
		tx:             tx,
	}
}

// CreateBannerWithEntities используется в случае создания уникального баннера
func (s *Service) CreateBannerWithEntities(ctx context.Context, banner entity.Banner) (int, error) {

	requestTags := make([]entity.Tag, 0, len(banner.TagIds))
	tagsToCreate := make([]entity.Tag, 0)
	for _, tagID := range banner.TagIds {
		tag := entity.Tag{ID: tagID}
		requestTags = append(requestTags, tag)
	}

	existingTags, err := s.tagService.FindExistingTags(ctx, requestTags)
	if err != nil {
		return 0, err
	}
	for _, tag := range requestTags {
		if _, ok := existingTags[tag]; !ok {
			tagsToCreate = append(tagsToCreate, tag)
		}
	}

	// todo: start tx
	var bannerID int
	err = s.tx.WithTx(ctx, func(ctx context.Context) error {
		var txErr error

		if len(tagsToCreate) > 0 {
			_, txErr = s.tagService.CreateTags(ctx, tagsToCreate)
			if txErr != nil {
				return txErr
			}
		}
		featID, txErr := s.featureService.FindOrCreateFeature(ctx, entity.Feature{ID: pointy.IntValue(banner.FeatureId, 0)})
		if txErr != nil {
			return txErr
		}

		//todo: improve converting
		bannerIsActive := pointy.BoolValue(banner.IsActive, true)
		banner = entity.Banner{
			FeatureId: pointy.Int(featID),
			TagIds:    banner.TagIds,
			IsActive:  &bannerIsActive,
			Content:   banner.Content,
		}
		bannerID, txErr = s.storage.CreateBanner(ctx, banner)
		if txErr != nil {
			return txErr
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return bannerID, nil
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
