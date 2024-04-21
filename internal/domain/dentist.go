package domain

type Dentist struct {
	IdDentist          int    `json:"id_dentist"`
	Name               string `json:"name"`
	Lastname           string `json:"lastname"`
	RegistrationNumber string `json:"registration_number"`
}
