package appointment

import (
	"context"
	"errors"
	"log"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dentist domain.Appointment) (domain.Appointment, error)
	GetAll(ctx context.Context) ([]domain.Appointment, error)
	GetByID(ctx context.Context, id int) (domain.Appointment, error)
	Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, requestAppointment domain.RequestAppointmentDateOnly, id int) (domain.Appointment, error)
}

// estructura del servicio, donde inyectamos el repositorio
type service struct {
	repository Repository
}

// constructor del service de appointment
func NewServiceAppointment(repository Repository) Service {
	return &service{repository: repository}
}

// Create ....
func (s *service) Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.repository.Create(ctx, appointment)
	if err != nil {
		log.Println("[AppointmentService][Create] error creating Appointment", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// GetAll ...apunta al service, misma firma
func (s *service) GetAll(ctx context.Context) ([]domain.Appointment, error) {
	listAppointments, err := s.repository.GetAll(ctx) //trae lista o error, llama al service, al rep que está adentro, implememta el get all del repository
	if err != nil {                                   //le pasa el contexto
		log.Println("[AppointmentService][GetAll] error getting all Appointment", err) //loguea el error
		return []domain.Appointment{}, err                                             //retorna el slice vacio y el error TODO ELLOGUEO DE ERRORES ES EN SERVICE
	}

	return listAppointments, nil //si no hay errores devuelve la lista de dentists y el error nulo
}

// GetByID ....
func (s *service) GetByID(ctx context.Context, id int) (domain.Appointment, error) {
	appointment, err := s.repository.GetByID(ctx, id) //le pasa el contexto y el id que buscamos
	if err != nil {
		log.Println("[AppointmentService][GetByID] error getting Appointment by ID", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// Update ...misma firma.
func (s *service) Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	appointment, err := s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentService][Update] error updating appointment by ID", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// Patch ... cambio de turno por fecha
func (s *service) Patch(ctx context.Context, requestAppointment domain.RequestAppointmentDateOnly, id int) (domain.Appointment, error) {
	foundAppointment, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Appointment{}, err
	}

	booleanAppointment, err := s.repository.Patch(ctx, requestAppointment, id)
	if err != nil || !booleanAppointment {
		log.Println("Error log in turn service. Patch Method", err.Error())
		return domain.Appointment{}, errors.New("Unexpected error")
	}

	foundAppointment.Datetime = requestAppointment.Datetime

	return foundAppointment, nil
}

// Delete ...
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[AppointmentService][Delete] error deleting Appointment by ID", err)
		return err
	}

	return nil
}
