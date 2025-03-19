package dto

import "time"

type GenericResponse struct {
	Rc      int         `json:"rc"`
	Mensaje string      `json:"msg"`
	Fecha   string      `json:"date"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(code int, msg string, obj interface{}) *GenericResponse {
	return &GenericResponse{
		Rc:      code,
		Mensaje: msg,
		Fecha:   time.Now().Format(time.RFC3339),
		Data:    obj,
	}
}
