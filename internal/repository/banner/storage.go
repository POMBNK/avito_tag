package banner

import (
	"context"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/pkg/client/postgres"
)

type Storage struct {
	conn postgres.Client
}

func New(conn postgres.Client) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) CreateBanner(ctx context.Context, banner entity.Banner) (int, error) {
	jsonb, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, err
	}
	insertBuilder := sq.Insert("banners").PlaceholderFormat(sq.Dollar).
		Columns("content", "feature_id", "is_active").
		Values(jsonb, banner.FeatureId, banner.IsActive).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var bannerID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&bannerID)
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}

func (s *Storage) AllBanners() {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) UpdateBanner() {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DeleteBanner() {
	//TODO implement me
	panic("implement me")
}
