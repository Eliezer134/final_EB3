package patient

import (
	"context"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, name domain.Patient, id int) (domain.Patient, error)
}
