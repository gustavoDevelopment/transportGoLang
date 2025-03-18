package dto

type User struct {
	DocumentType   string `json:"document_type"`
	DocumentNumber string `json:"document_number"`
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	FirstLastName  string `json:"first_last_name"`
	SecondLastName string `json:"second_last_name"`
}

type GenericResponse struct {
	Rc      float32 `json:"rc"`
	Mensaje string  `json:"msg"`
	Fecha   string  `json:"date"`
	Data    any     `json:"data,omitempty"`
}
