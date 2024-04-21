package domain

import (
	"time"
)

type Appointment struct {
	IdAppointment int       `json:"id_appointment"`
	Description   string    `json:"description"`
	Dentist       Dentist   `json:"dentist" binding:"required" `
	Patient       Patient   `json:"patient" binding:"required" `
	Datetime      time.Time `json:"datetime" `
}

type RequestAppointment struct {
	Description string    `json:"description" `
	Id_patient  int       `json:"id_patient" binding:"required" `
	Id_dentist  int       `json:"id_dentist" binding:"required" `
	Datetime    time.Time `json:"datetime" binding:"required"`
}

type RequestAppointmentByDniAndRegistrationNumber struct {
	Description        string    `json:"description" `
	Dni                string    `json:"dni" binding:"required" `
	RegistrationNamber string    `json:"registrationNamber" binding:"required" `
	Datetime           time.Time `json:"datetime" binding:"required"`
}

type RequestAppointmentDateOnly struct {
	IdAppointment int       `json:"id_appointment"`
	Datetime      time.Time `json:"datetime" binding:"required"`
}
