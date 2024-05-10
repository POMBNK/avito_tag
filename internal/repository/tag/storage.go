package tag

import (
	"context"
	"fmt"
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

func (s *Storage) CreateTag(ctx context.Context, tag entity.Tag) (int, error) {
	insertBuilder := sq.Insert("tags").PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(tag.Name).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var tagID int
	err = s.conn.QueryRow(ctx, query, args...).Scan(&tagID)
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

func (s *Storage) CreateTags(ctx context.Context, tags []entity.Tag) ([]int, error) {
	insertBuilder := sq.Insert("tags").Columns("id", "name").PlaceholderFormat(sq.Dollar)

	for _, tag := range tags {
		tag.Name = fmt.Sprintf("tag-%d", tag.ID)
		insertBuilder = insertBuilder.Values(tag.ID, tag.Name)
	}
	insertBuilder = insertBuilder.Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	sqlRows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var tagID int
	tagIDs := make([]int, 0, len(tags))
	for sqlRows.Next() {
		err = sqlRows.Scan(&tagID)
		if err != nil {
			return nil, err
		}

		tagIDs = append(tagIDs, tagID)
	}

	if sqlRows.Err() != nil {
		return nil, sqlRows.Err()
	}

	return tagIDs, nil

}

func (s *Storage) FindExistingTags(ctx context.Context, tags []entity.Tag) (map[entity.Tag]struct{}, error) {

	tagIDs := make([]int, 0, len(tags))
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	selectBuilder := sq.Select("id").PlaceholderFormat(sq.Dollar).
		From("tags").
		Where(sq.Eq{"id": tagIDs})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	sqlRows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var tagID int
	existingTags := make(map[entity.Tag]struct{}, len(tags))
	for sqlRows.Next() {
		err = sqlRows.Scan(&tagID)
		if err != nil {
			return nil, err
		}

		existingTags[entity.Tag{ID: tagID}] = struct{}{}
	}

	return existingTags, nil
}
