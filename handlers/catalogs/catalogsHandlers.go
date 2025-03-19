package catalogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restMux/database"
	"restMux/dto"
	"restMux/logger"
	"restMux/model/catalog"
	"time"
)

func HandleCatalogRequest(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("HandleCatalogRequest")
	rs.Header().Set("Content-Type", "application/json")
	queryType := rq.URL.Query().Get("type")
	if queryType == "" {
		currentTime := time.Now().Format(time.RFC3339)
		genRes := dto.GenericResponse{Rc: -2000, Mensaje: "Falta el parámetro 'type'", Fecha: currentTime}
		response, _ := json.Marshal(genRes)
		rs.WriteHeader(http.StatusBadRequest)
		rs.Write(response)
		return
	}

	switch queryType {
	case "eps":

		doOnGetEPS(rs, rq)

	case "documentType":
		doOnGetDocumentType(rs, rq)
	case "deparment":
		doOnGetDepartments(rs, rq)
	case "city":
		doOnGetCities(rs, rq)
	case "insuranceCompany":
		doOnGetInsurers(rs, rq)
	case "truckBrand":
		doOnGetTruckBrands(rs, rq)
	default:
		currentTime := time.Now().Format(time.RFC3339)
		genRes := dto.GenericResponse{Rc: -2000, Mensaje: "Tipo de catálogo: [" + queryType + "] no válido", Fecha: currentTime}
		response, _ := json.Marshal(genRes)
		rs.WriteHeader(http.StatusBadRequest)
		rs.Write(response)
		return
	}
}

func doOnGetEPS(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetEPS")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	data := catalog.EPSs{}
	database.Database.Find(&data)
	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catalogo de EPS", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetDocumentType(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetDocumentType")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	data := catalog.DocumentTypes{}
	database.Database.Order("id ASC").Find(&data)
	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catalogo de DocumentType", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetDepartments(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetDepartments")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	data := catalog.Departments{}
	database.Database.Find(&data)

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de departamentos", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetCities(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetCities")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	data := catalog.Citys{}
	database.Database.Find(&data)

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de ciudades", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetInsurers(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetInsurers")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	data := catalog.InsuranceCompanys{}
	database.Database.Find(&data)
	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de aseguradoras", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetTruckBrands(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetTruckBrands")
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	data := catalog.TruckBrands{}
	database.Database.Find(&data)

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: data}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de marcas de camiones", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}
