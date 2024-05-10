package feature

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/POMBNK/avitotag/internal/entity"
	"github.com/POMBNK/avitotag/internal/pkg/apierror"
	"github.com/POMBNK/avitotag/internal/pkg/client/postgres"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn postgres.Client
}

func New(conn postgres.Client) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) FindOrCreateFeature(ctx context.Context, feature entity.Feature) (int, error) {

	// todo: вынести в FindByID
	selectBuilder := sq.Select("id").From("features").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": feature.ID})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var featureID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&featureID)
	if err == nil {
		// todo: check if err sql.ErrNoRows
		return 0, nil
	}

	insertBuilder := sq.Insert("features").PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(feature.Name).
		Suffix("RETURNING id")

	query, args, err = insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	err = s.conn.QueryRow(ctx, query, args...).Scan(&featureID)

	if err != nil {
		return 0, err
	}

	return featureID, nil
}

func (s *Storage) FindByID(ctx context.Context, featureID int) (entity.Feature, error) {
	selectBuilder := sq.Select("id").From("features").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": featureID})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return entity.Feature{}, err
	}

	var foundFeatureID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&foundFeatureID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Feature{}, apierror.ErrNotFound
		}
		return entity.Feature{}, err
	}

	return entity.Feature{ID: foundFeatureID}, nil
}

func (s *Storage) CreateFeature(ctx context.Context, feature entity.Feature) (int, error) {

	feature.Name = fmt.Sprintf("feature-%d", feature.ID)

	insertBuilder := sq.Insert("features").PlaceholderFormat(sq.Dollar).
		Columns("id", "name").
		Values(feature.ID, feature.Name).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var featureID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&featureID)
	if err != nil {
		return 0, err
	}

	return featureID, nil
}
