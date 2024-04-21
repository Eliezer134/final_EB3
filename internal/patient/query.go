package patient

var (
	QueryInsertPatient = `INSERT INTO odontologia.patients(name, lastname, address, dni, registration_date)
    VALUES(?,?,?,?,?)`
	QueryGetPatientByID = `SELECT id_patient, name, lastname, address, dni, registration_date
    FROM odontologia.patients WHERE id_patient = ?`
	QueryDeletePatientByID = `DELETE FROM odontologia.patients WHERE id_patient = ?`
	QueryUpdatePatientByID = `UPDATE odontologia.patients SET name = ?, lastname= ?, address = ?, dni= ?, registration_date= ? WHERE id_patient = ?`
	QueryPatchPatientName  = `UPDATE odontologia.patients SET name = ? WHERE id_patient = ?`
	QueryGetAll            = `SELECT * FROM odontologia.patients`
)
