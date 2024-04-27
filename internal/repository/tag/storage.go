package tag

import "github.com/POMBNK/avitotag/internal/pkg/client/postgres"

type Storage struct {
	conn postgres.Client
}

func New(conn postgres.Client) *Storage {
	return &Storage{conn: conn}
}

func (s *Storage) CreateBanner() {
	//TODO implement me
	panic("implement me")
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
