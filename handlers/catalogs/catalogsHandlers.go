package catalogs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restMux/constanst"
	"restMux/database"
	"restMux/dto"
	"restMux/logger"
	"restMux/model/catalog"
)

func HandleCatalogRequest(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("HandleCatalogRequest")
	rs.Header().Set("Content-Type", "application/json")
	queryType := rq.URL.Query().Get("type")
	if queryType == "" {
		genRes := dto.NewResponse(constanst.PROCESS_ERROR, "Falta el parámetro 'type'", nil)
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
		genRes := dto.NewResponse(-2000, "Tipo de catálogo: ["+queryType+"] no válido", nil)
		response, _ := json.Marshal(genRes)
		rs.WriteHeader(http.StatusBadRequest)
		rs.Write(response)
		return
	}
}

func doOnGetEPS(rs http.ResponseWriter, rq *http.Request) {
	logger.Log.Info("doOnGetEPS")
	rs.Header().Set("Content-Type", "application/json")
	data := catalog.EPSs{}
	database.Database.Find(&data)
	var genRes *dto.GenericResponse
	genRes = dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)
	response, err := json.Marshal(genRes)
	logger.Log.Info(string(response))
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
	data := catalog.DocumentTypes{}
	database.Database.Order("id ASC").Find(&data)
	genRes := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)

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

	data := catalog.Departments{}
	database.Database.Find(&data)

	genRes := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)

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

	data := catalog.Citys{}
	database.Database.Find(&data)

	genRes := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)
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

	data := catalog.InsuranceCompanys{}
	database.Database.Find(&data)
	genRes := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)

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

	data := catalog.TruckBrands{}
	database.Database.Find(&data)
	genRes := dto.NewResponse(constanst.PROCESS_OK, constanst.PROCESS_OK_MSG, data)
	response, err := json.Marshal(genRes)
	if err != nil {
		http.Error(rs, "Error al obtener el catálogo de marcas de camiones", http.StatusInternalServerError)
		return
	}

	rs.WriteHeader(http.StatusOK)
	fmt.Fprintln(rs, string(response))
}
