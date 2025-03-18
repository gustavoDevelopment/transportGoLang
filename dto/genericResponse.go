package dto

type GenericResponse struct {
	Rc      float32 `json:"rc"`
	Mensaje string  `json:"msg"`
	Fecha   string  `json:"date"`
	Data    any     `json:"data,omitempty"`
}
