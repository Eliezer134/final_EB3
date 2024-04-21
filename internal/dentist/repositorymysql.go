package dentist

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

// NewMemoryRepository
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create
func (r *repositorymysql) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentist)
	if err != nil {
		return domain.Dentist{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.Lastname,
		dentist.RegistrationNumber,
	)

	if err != nil {
		return domain.Dentist{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, ErrLastInsertedId
	}

	dentist.IdDentist = int(lastId)

	return dentist, nil

}

// get all dentist
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Dentist, error) {

	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Dentist{}, err
	}

	defer rows.Close()

	var dentists []domain.Dentist

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(
			&dentist.IdDentist,
			&dentist.Name,
			&dentist.Lastname,
			&dentist.RegistrationNumber,
		)
		if err != nil {
			return []domain.Dentist{}, err
		}

		dentists = append(dentists, dentist)
	}

	if err := rows.Err(); err != nil {
		return []domain.Dentist{}, err
	}

	return dentists, nil
}

// GetByID
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistByID, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.IdDentist,
		&dentist.Name,
		&dentist.Lastname,
		&dentist.RegistrationNumber,
	)

	if err != nil {
		// Verifica si el error es sql.ErrNoRows y retorna ErrNotFound en ese caso
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Dentist{}, ErrNotFound
		}

		// Si el error no es sql.ErrNoRows, retorna el error original
		return domain.Dentist{}, err
	}

	// Si no hay error, devuelve la estructura del dentista
	return dentist, nil
}

// Update
func (r *repositorymysql) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	// Prepare the update statement
	stmt, err := r.db.Prepare(QueryUpdateDentistByID)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer stmt.Close()

	// Bind the values and execute the update statement
	_, err = stmt.Exec(
		dentist.IdDentist,
		dentist.Name,
		dentist.Lastname,
		dentist.RegistrationNumber,
	)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, err

}

// Patch
func (r *repositorymysql) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {

	stmt, err := r.db.Prepare(QueryPatchDentistName)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer stmt.Close()

	// Bind the values and execute the patch statement
	result, err := stmt.Exec(dentist.Name, id)
	if err != nil {
		return domain.Dentist{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}

// Delete ...
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentistByID, id)
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
