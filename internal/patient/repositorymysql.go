package patient

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
	ErrNotFound         = errors.New("not found")
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create ....
func (r *repositorymysql) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return domain.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Lastname,
		patient.Address,
		patient.Dni,
		patient.RegistrationDate,
	)

	if err != nil {
		return domain.Patient{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Patient{}, ErrLastInsertedId
	}

	patient.IdPatient = int(lastId)

	return patient, nil

}

// GetAll...
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Patient, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Patient{}, err
	}

	defer rows.Close()

	var patients []domain.Patient

	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(
			&patient.IdPatient,
			&patient.Name,
			&patient.Lastname,
			&patient.Address,
			&patient.Dni,
			&patient.RegistrationDate,
		)
		if err != nil {
			return []domain.Patient{}, err
		}

		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return []domain.Patient{}, err
	}

	return patients, nil
}

// GetByID .....
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	row := r.db.QueryRow(QueryGetPatientByID, id)

	var patient domain.Patient
	err := row.Scan(

		&patient.IdPatient,
		&patient.Name,
		&patient.Lastname,
		&patient.Address,
		&patient.Dni,
		&patient.RegistrationDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Patient{}, ErrNotFound
		}
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update ...
func (r *repositorymysql) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	stmt, err := r.db.Prepare(QueryUpdatePatientByID)
	if err != nil {
		return domain.Patient{}, err
	}

	defer stmt.Close()

	// Bind the values and execute the update statement
	_, err = stmt.Exec(

		patient.Name,
		patient.Lastname,
		patient.Address,
		patient.Dni,
		patient.RegistrationDate,
		patient.IdPatient,
	)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, err

}

//patch

func (r *repositorymysql) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	// Prepare the update statement
	stmt, err := r.db.Prepare(QueryPatchPatientName)
	if err != nil {
		return domain.Patient{}, err
	}

	defer stmt.Close()

	// Bind the values and execute the patch statement
	result, err := stmt.Exec(patient.Name, id)
	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

// Delete ...
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePatientByID, id)
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
