package patient

import (
	"context"
	"log"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

type Service interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
}

// crea la estructura service que va a tener el repositorio donde inyectamos la dependencia, estamos inyectando el repositorio dentro de la estruct
type service struct {
	repository Repository
}

// constructor de un service, recibe un repo de parametro y devuelve algo que cumpla con la interfaz de service
func NewServicePatient(repository Repository) Service {
	return &service{repository: repository}
}

// Create ....
func (s *service) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[PatientService][Create] error creating patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// GetAll ...apunta al service, misma firma
func (s *service) GetAll(ctx context.Context) ([]domain.Patient, error) {
	listPatients, err := s.repository.GetAll(ctx) //trae lista o error, llama al service, al rep que está adentro, implememta el get all del repository
	if err != nil {                               //le pasa el contexto
		log.Println("[PatientsService][GetAll] error getting all patients", err) //loguea el error
		return []domain.Patient{}, err                                           //retorna el slice vacio y el error TODO ELLOGUEO DE ERRORES ES EN SERVICE
	}

	return listPatients, nil //si no hay errores devuelve la lista de dentists y el error nulo
}

// GetByID ....
func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id) //le pasa el contexto y el id que buscamos
	if err != nil {
		log.Println("[PatientsService][GetByID] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update ...misma firma.
func (s *service) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Patch ...
func (s *service) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	// Llamar al método de parche del repositorio con el paciente y el ID
	_, err := s.repository.Patch(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error patching patient by ID", err)
		return domain.Patient{}, err
	}

	// Obtener el paciente actualizado con todos los campos desde la base de datos
	patchedPatient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error getting patched patient by ID", err)
		return domain.Patient{}, err
	}

	return patchedPatient, nil
}

// Delete ...
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Delete] error deleting patient by ID", err)
		return err
	}

	return nil
}
