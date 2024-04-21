package dentist

import (
	"context"
	"log"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetAll(ctx context.Context) ([]domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewServiceDentist(repository Repository) Service {
	return &service{repository: repository}
}

// Create
func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistService][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// GetAll
func (s *service) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	listDentists, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[DentistsService][GetAll] error getting all dentists", err)
		return []domain.Dentist{}, err
	}

	return listDentists, nil
}

// GetByID
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][GetByID] error getting dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update
func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistsService][Update] error updating product by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}
func (s *service) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {

	_, err := s.repository.Patch(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistsService][Patch] error patching dentist by ID", err)
		return domain.Dentist{}, err
	}
	patchedDentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Patch] error getting patched dentist by ID", err)
		return domain.Dentist{}, err
	}

	return patchedDentist, nil
}

// Delete
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Delete] error deleting dentist by ID", err)
		return err
	}

	return nil
}
