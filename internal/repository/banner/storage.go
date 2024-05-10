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

	insertLinkBuilder := sq.Insert("banner_tags").PlaceholderFormat(sq.Dollar).
		Columns("banner_id", "tag_id")

	for _, tagID := range banner.TagIds {
		insertLinkBuilder = insertLinkBuilder.
			Values(bannerID, tagID)
	}

	query, args, err = insertLinkBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	_, err = s.conn.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return bannerID, nil
}

func (s *Storage) AllBanners(ctx context.Context, params entity.BannerSearchParams) ([]entity.Banner, error) {
	selectBuilder := sq.Select(
		"b.id",
		"b.name",
		"b.feature_id",
		"b.created_at",
		"b.content",
		"b.is_active",
		"(SELECT array_agg(tag_id) FROM banner_tags bt WHERE bt.banner_id = b.id ) as tag_ids").
		From("banners b").PlaceholderFormat(sq.Dollar)

	if params.FeatureId != nil {
		selectBuilder = selectBuilder.Where(sq.Eq{"b.feature_id": params.FeatureId})
	}
	if params.TagId != nil {
		selectBuilder = selectBuilder.Join("banner_tags bt ON b.id = bt.banner_id").
			Where(sq.Eq{"bt.tag_id": params.TagId})
	}

	if params.Limit != nil {
		selectBuilder = selectBuilder.Limit(uint64(*params.Limit))
	} else {
		selectBuilder = selectBuilder.Limit(10)
	}
	if params.Offset != nil {
		selectBuilder = selectBuilder.Offset(uint64(*params.Offset))
	} else {
		selectBuilder = selectBuilder.Offset(0)
	}

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var banners []entity.Banner
	rows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var banner entity.Banner
		err = rows.Scan(&banner.ID, &banner.Name, &banner.FeatureId, &banner.CreatedAt, &banner.Content, &banner.IsActive, &banner.TagIds)
		if err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}

func (s *Storage) UpdateBanner() {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DeleteBanner() {
	//TODO implement me
	panic("implement me")
}

func BannerByFeatureIDAndTagIDs() {
	/*
		SELECT b.id, b.name
		FROM banners b
		         JOIN banner_tags bt ON b.id = bt.banner_id
		         JOIN tags t ON bt.tag_id = t.id
		WHERE b.feature_id = 2
		  AND bt.tag_id IN (1,2,3) -- Перечислите идентификаторы тегов здесь
		GROUP BY b.id, b.name
	*/
}
