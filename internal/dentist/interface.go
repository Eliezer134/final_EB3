package dentist

import (
	"context"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetAll(ctx context.Context) ([]domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Patch(ctx context.Context, name domain.Dentist, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
}
