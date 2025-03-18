package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restMux/dto"
	"time"

	"github.com/gorilla/mux"
)

func DoOnEjemploGet(rs http.ResponseWriter, rq *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	response := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(rs, "Error al procesar la respuesta", http.StatusInternalServerError)
		return
	}
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(output))
}

func DoOnEjemploGetWithParameter(rs http.ResponseWriter, rq *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	vars := mux.Vars(rq)
	response := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: vars["id"]}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(rs, "Error al procesar la respuesta", http.StatusInternalServerError)
		return
	}
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(output))
}

func DoOnEjemploPost(rs http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rs, "doOnEjemploPost")
}

func DoOnEjemploPostWithBody(rs http.ResponseWriter, rq *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	rs.Header().Set("Content-Type", "application/json")
	auth := rq.Header.Get("Authorization")
	if auth == "" {
		output, _ := json.Marshal(dto.GenericResponse{Rc: -2000, Mensaje: "not Authorization request", Fecha: currentTime})
		rs.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(rs, string(output))
		return
	}

	var user dto.User
	err := json.NewDecoder(rq.Body).Decode(&user)
	if err != nil {
		http.Error(rs, "Error al procesar el body", http.StatusInternalServerError)
		return
	}
	response := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: user.FirstName + " " + user.SecondName}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(rs, "Error al procesar la respuesta", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(output))
}

func DoOnEjemploPut(rs http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	fmt.Fprintln(rs, "doOnEjemploPut |id: "+vars["id"])
}

func DoOnEjemploDelete(rs http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	fmt.Fprintln(rs, "doOnEjemploDelete |id: "+vars["id"])
}

func DoOnEjemploGetQueryString(rs http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rs, "doOnEjemploGetQueryString |id: "+rq.URL.Query().Get("id")+"| name: "+rq.URL.Query().Get("name"))
}
