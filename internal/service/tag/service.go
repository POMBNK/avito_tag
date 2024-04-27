package tag

import "github.com/POMBNK/avitotag/internal/repository"

type Service struct {
	storage repository.TagRepository
}

func New(storage repository.TagRepository) *Service {
	return &Service{storage: storage}
}
