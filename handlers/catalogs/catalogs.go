package catalogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restMux/dto"
	"restMux/model/catalog"
	"time"
)

func HandleCatalogRequest(rs http.ResponseWriter, rq *http.Request) {
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
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	var EPSList = []catalog.EPS{
		{ID: "1", Code: "EPS001", Name: "Sanitas EPS"},
		{ID: "2", Code: "EPS002", Name: "Nueva EPS"},
		{ID: "3", Code: "EPS003", Name: "Compensar EPS"},
		{ID: "4", Code: "EPS004", Name: "Sura EPS"},
		{ID: "5", Code: "EPS005", Name: "Coomeva EPS"},
		{ID: "6", Code: "EPS006", Name: "Famisanar EPS"},
		{ID: "7", Code: "EPS007", Name: "Salud Total EPS"},
		{ID: "8", Code: "EPS008", Name: "Mutual Ser EPS"},
		{ID: "9", Code: "EPS009", Name: "Medimás EPS"},
		{ID: "10", Code: "EPS010", Name: "Cafesalud EPS"},
		{ID: "11", Code: "EPS011", Name: "Emssanar EPS"},
		{ID: "12", Code: "EPS012", Name: "Asmet Salud EPS"},
		{ID: "13", Code: "EPS013", Name: "Ecoopsos EPS"},
		{ID: "14", Code: "EPS014", Name: "Capital Salud EPS"},
		{ID: "15", Code: "EPS015", Name: "SOS EPS"},
	}

	//data := catalog.EPSs{}
	//database.Database.Find(&data)
	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: EPSList}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catalogo de EPS", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetDocumentType(rs http.ResponseWriter, rq *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	var documentTypes = []catalog.DocumentType{
		{ID: "1", Code: "CC", Name: "Cédula de Ciudadanía"},
		{ID: "2", Code: "TI", Name: "Tarjeta de Identidad"},
		{ID: "3", Code: "CE", Name: "Cédula de Extranjería"},
		{ID: "4", Code: "NIT", Name: "Número de Identificación Tributaria"},
		{ID: "5", Code: "PP", Name: "Pasaporte"},
	}

	//data := catalog.EPSs{}
	//database.Database.Find(&data)
	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: documentTypes}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catalogo de DocumentType", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetDepartments(rs http.ResponseWriter, rq *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	var departments = []catalog.Department{
		{ID: "1", Code: "05", Name: "Antioquia"},
		{ID: "2", Code: "08", Name: "Atlántico"},
		{ID: "3", Code: "11", Name: "Bogotá D.C."},
	}

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: departments}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de departamentos", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetCities(rs http.ResponseWriter, rq *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	var cities = []catalog.City{
		{ID: "1", Code: "11001", Name: "Bogotá"},
		{ID: "2", Code: "05001", Name: "Medellín"},
		{ID: "3", Code: "76001", Name: "Cali"},
	}

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: cities}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de ciudades", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetInsurers(rs http.ResponseWriter, rq *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	var insurers = []catalog.InsuranceCompany{
		{ID: "1", Code: "001", Name: "Sura"},
		{ID: "2", Code: "002", Name: "Bolívar"},
		{ID: "3", Code: "003", Name: "Colpatria"},
	}

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: insurers}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de aseguradoras", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}

func doOnGetTruckBrands(rs http.ResponseWriter, rq *http.Request) {
	rs.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)

	var brands = []catalog.TruckBrand{
		{ID: "1", Code: "VOL", Name: "Volvo"},
		{ID: "2", Code: "SCAN", Name: "Scania"},
		{ID: "3", Code: "KEN", Name: "Kenworth"},
	}

	genRes := dto.GenericResponse{Rc: 0, Mensaje: "process Ok", Fecha: currentTime, Data: brands}

	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de marcas de camiones", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}
