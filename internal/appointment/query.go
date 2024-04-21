package appointment

var (
	// internal\appointment\query.go

	QueryGetAppointmentByID = `SELECT a.id_appointment, a.description, d.id_dentist, d.name, d.lastname, d.registration_number,
	p.id_patient, p.name, p.lastname, p.address, p.dni, p.registration_date, a.datetime
	FROM odontologia.appointments a
	JOIN odontologia.dentists d ON d.registration_number = a.registration_number
	JOIN odontologia.patients p ON p.dni = a.dni
	WHERE a.id_appointment = ?`
	QueryGetAllByPatientDni = `SELECT a.id, a.description, a.datetime, d.id, d.name, d.lastname, d.registration_number, 
	p.id, p.name, p.lastname, p.address, p.dni, p.registration_date
	FROM odontologia.appointments a
	JOIN odontologia.dentists d ON d.registration_number = a.registration_number
	JOIN odontologia.patients p ON p.dni = a.dni
	WHERE p.dni = ?`
	QueryGetAllByDentistRegistrationNumber = `SELECT a.id, a.description, a.datetime, d.id, d.name, d.lastname, d.registration_number, 
	p.id, p.name, p.lastname, p.address, p.dni, p.registration_date
	FROM odontologia.appointments a
	JOIN odontologia.dentists d ON d.registration_number = a.registration_number
	JOIN odontologia.patients p ON p.dni = a.dni
	WHERE d.registration_number = ?`
	QueryGetAll = `SELECT a.id_appointment, a.description, d.id_dentist, d.name, d.lastname, d.registration_number,
	p.id_patient, p.name, p.lastname, p.address, p.dni, p.registration_date, a.datetime
	FROM odontologia.appointments a
	JOIN odontologia.dentists d ON d.registration_number = a.registration_number
	JOIN odontologia.patients p ON p.dni = a.dni`

	QueryInsertAppointment     = `INSERT INTO odontologia.appointments (description, registration_number, dni, datetime) VALUES (?,?,?,?)`
	QueryDeleteAppointmentByID = `DELETE FROM odontologia.appointments WHERE id_appointment = ?`
	QueryUpdateAppointment     = `UPDATE odontologia.appointments SET datetime = ?, description= ?, registration_number= ?, dni= ? WHERE id_appointment = ?`
	QueryPatchDateAppointment  = `PATCH odontologia.appointments SET datetime = ? WHERE id_appointment = ?`
)
