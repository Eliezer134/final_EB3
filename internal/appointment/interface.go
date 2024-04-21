package appointment

import (
	"context"
	"errors"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

var (
	ErrEmptyList = errors.New("The list is empty")
	ErrNotFound  = errors.New("Turn not found")
	ErrStatement = errors.New("Error preparing statement")
	ErrExec      = errors.New("Error exect statement")
	ErrLastID    = errors.New("Error getting last id")
)

type Repository interface {
	Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error)
	GetAll(ctx context.Context) ([]domain.Appointment, error)
	GetByID(ctx context.Context, id int) (domain.Appointment, error)
	Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
	Patch(ctx context.Context, requestAppointment domain.RequestAppointmentDateOnly, id int) (bool, error)
	Delete(ctx context.Context, id int) error
}
