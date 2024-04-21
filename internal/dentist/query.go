package dentist

var (
	QueryInsertDentist = `INSERT INTO odontologia.dentists(name, lastname, registration_number)
	VALUES(?,?,?)`
	QueryGetDentistByID = `SELECT id_dentist, name, lastname, registration_number
	FROM odontologia.dentists WHERE id_dentist = ?`
	QueryDeleteDentistByID              = `DELETE FROM odontologia.dentists WHERE id_dentist = ?`
	QueryUpdateDentistByID              = `UPDATE odontologia.dentists SET name = ?, lastname= ?, registration_number= ? WHERE id_dentist = ?`
	QueryGetDentistByRegistrationNumber = `SELECT id_dentist, name, lastname, registration_number
	FROM odontologia.dentists WHERE registration_number = ?`
	QueryPatchDentistName = `UPDATE odontologia.dentists SET name = ? WHERE id_dentist = ?`
	QueryGetAll           = `SELECT * FROM odontologia.dentists`
)
