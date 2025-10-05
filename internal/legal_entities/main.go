package legalentities

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func New(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetLegalEntities(ctx context.Context) ([]LegalEntity, error) {
	_ = ctx
	return s.repo.WithCounter().getAllLegalEntities()
}

func (s *Service) UpdateLegalEntityName(ctx context.Context, uid uuid.UUID, name string) error {
	_ = ctx
	return s.UpdateLegalEntity(uid, name)
}

func (s *Service) CreateLegalEntity(name string) (LegalEntity, error) {
	if name == "" {
		return LegalEntity{}, errors.New("name is required")
	}
	return s.repo.CreateLegalEntity(name)
}

func (s *Service) UpdateLegalEntity(uid uuid.UUID, name string) error {
	if uid == uuid.Nil {
		return errors.New("uuid is required")
	}
	if name == "" {
		return errors.New("name is required")
	}
	return s.repo.UpdateLegalEntity(uid, name)
}

func (s *Service) DeleteLegalEntity(uid uuid.UUID) error {
	if uid == uuid.Nil {
		return errors.New("uuid is required")
	}
	return s.repo.DeleteLegalEntity(uid)
}
