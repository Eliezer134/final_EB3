package appointment

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Eliezer134/final_EB3.git/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// GetByID returns a turn by the ID sent as parameter
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Appointment, error) {
	row := r.db.QueryRow(QueryGetAppointmentByID, id)

	var appointment domain.Appointment
	err := row.Scan(
		&appointment.IdAppointment,
		&appointment.Description,
		&appointment.Dentist.IdDentist,
		&appointment.Dentist.Name,
		&appointment.Dentist.Lastname,
		&appointment.Dentist.RegistrationNumber,
		&appointment.Patient.IdPatient,
		&appointment.Patient.Name,
		&appointment.Patient.Lastname,
		&appointment.Patient.Address,
		&appointment.Patient.Dni,
		&appointment.Patient.RegistrationDate,
		&appointment.Datetime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Appointment{}, ErrNotFound
		}
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (r *repositorymysql) GetAllByPatientDni(ctx context.Context, dni string) ([]domain.Appointment, error) {

	rows, err := r.db.Query(QueryGetAllByPatientDni, dni)
	if err != nil {
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	var appointments []domain.Appointment

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.IdAppointment,
			&appointment.Description,
			&appointment.Datetime,
			&appointment.Dentist.IdDentist,
			&appointment.Dentist.Name,
			&appointment.Dentist.Lastname,
			&appointment.Dentist.RegistrationNumber,
			&appointment.Patient.IdPatient,
			&appointment.Patient.Name,
			&appointment.Patient.Lastname,
			&appointment.Patient.Address,
			&appointment.Patient.Dni,
			&appointment.Patient.RegistrationDate,
		)
		if err != nil {
			return []domain.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return []domain.Appointment{}, err
	}

	return appointments, nil
}

// get all appointment
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Appointment, error) {

	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	var appointments []domain.Appointment

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.IdAppointment,
			&appointment.Description,
			&appointment.Dentist.IdDentist,
			&appointment.Dentist.Name,
			&appointment.Dentist.Lastname,
			&appointment.Dentist.RegistrationNumber,
			&appointment.Patient.IdPatient,
			&appointment.Patient.Name,
			&appointment.Patient.Lastname,
			&appointment.Patient.Address,
			&appointment.Patient.Dni,
			&appointment.Patient.RegistrationDate,
			&appointment.Datetime,
		)
		if err != nil {
			return []domain.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return []domain.Appointment{}, err
	}

	return appointments, nil
}

// Create creates a new Appointment.
func (r *repositorymysql) Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error) {
	statement, err := r.db.Prepare(QueryInsertAppointment)

	if err != nil {
		return domain.Appointment{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(

		appointment.Description,
		appointment.Dentist.RegistrationNumber,
		appointment.Patient.Dni,
		appointment.Datetime,
	)

	if err != nil {
		return domain.Appointment{}, ErrExecStatement
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Appointment{}, ErrLastInsertedId
	}

	appointment.IdAppointment = int(lastID)

	return appointment, nil
}

// DeleteByID deletes a turn by the ID sent as parameter
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteAppointmentByID, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}

func (r *repositorymysql) Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	// Preparar la sentencia de actualización
	stmt, err := r.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	defer stmt.Close()

	// Ejecutar la sentencia de actualización con los parámetros correspondientes
	_, err = stmt.Exec(
		appointment.Datetime,
		appointment.Description,
		appointment.Dentist.RegistrationNumber,
		appointment.Patient.Dni,
		appointment.IdAppointment, // ID de la cita que se va a actualizar
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	// Devolver la cita actualizada
	return appointment, nil
}

func (r *repositorymysql) Patch(ctx context.Context, requestAppointment domain.RequestAppointmentDateOnly, id int) (bool, error) {
	// Prepare the update statement
	stmt, err := r.db.Prepare(QueryPatchDateAppointment)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	// Bind the values and execute the update statement
	_, err = stmt.Exec(
		requestAppointment.Datetime,
		requestAppointment.IdAppointment,
	)
	if err != nil {
		return false, err
	}

	return true, err
}

func (r *repositorymysql) GetAllByDentistRegistrationNumber(ctx context.Context, registrationNumber string) ([]domain.Appointment, error) {

	rows, err := r.db.Query(QueryGetAllByDentistRegistrationNumber, registrationNumber)
	if err != nil {
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	var appointments []domain.Appointment

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.IdAppointment,
			&appointment.Description,
			&appointment.Datetime,
			&appointment.Dentist.IdDentist,
			&appointment.Dentist.Name,
			&appointment.Dentist.Lastname,
			&appointment.Dentist.RegistrationNumber,
			&appointment.Patient.IdPatient,
			&appointment.Patient.Name,
			&appointment.Patient.Lastname,
			&appointment.Patient.Address,
			&appointment.Patient.Dni,
			&appointment.Patient.RegistrationDate,
		)
		if err != nil {
			return []domain.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return []domain.Appointment{}, err
	}

	return appointments, nil
}
