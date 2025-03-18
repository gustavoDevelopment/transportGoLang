package model

type Person struct {
	DocumentType   string `json:"document_type"`              // Tipo de documento (Ej: DNI, PASAPORTE)
	DocumentNumber string `json:"document_number"`            // Número de documento
	FirstName      string `json:"first_name"`                 // Primer nombre
	SecondName     string `json:"second_name,omitempty"`      // Segundo nombre (opcional)
	FirstLastName  string `json:"first_last_name"`            // Primer apellido
	SecondLastName string `json:"second_last_name,omitempty"` // Segundo apellido (opcional)
	DateOfBirth    string `json:"date_of_birth,omitempty"`    // Fecha de nacimiento (Formato: YYYY-MM-DD)
	Gender         string `json:"gender,omitempty"`           // Género (Ej: Male, Female, Other)
	Nationality    string `json:"nationality,omitempty"`      // Nacionalidad
	PhoneNumber    string `json:"phone_number,omitempty"`     // Número de teléfono
	Email          string `json:"email,omitempty"`            // Correo electrónico
	Address        string `json:"address,omitempty"`          // Dirección
}
