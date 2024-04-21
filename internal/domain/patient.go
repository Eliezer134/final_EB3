package domain

type Patient struct {
	IdPatient        int    `json:"id_patient"`
	Name             string `json:"name"`
	Lastname         string `json:"lastname"`
	Address          string `json:"address"`
	Dni              string `json:"dni"`
	RegistrationDate string `json:"registration_date"`
}
