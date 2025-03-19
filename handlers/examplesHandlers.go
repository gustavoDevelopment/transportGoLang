package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restMux/constanst"
	"restMux/dto"
	"restMux/model"

	"github.com/gorilla/mux"
)

func DoOnEjemploGet(rs http.ResponseWriter, rq *http.Request) {
	response := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, nil)

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
	vars := mux.Vars(rq)
	response := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, vars["id"])
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
	rs.Header().Set("Content-Type", "application/json")
	auth := rq.Header.Get("Authorization")
	if auth == "" {
		output, _ := json.Marshal(dto.NewResponse(constanst.PROCESS_ERROR_AUTH, constanst.PROCESS_ERROR_AUTH_MSG, nil))
		rs.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(rs, string(output))
		return
	}

	var user model.Person
	err := json.NewDecoder(rq.Body).Decode(&user)
	if err != nil {
		http.Error(rs, "Error al procesar el body", http.StatusInternalServerError)
		return
	}
	response := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, user)
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
