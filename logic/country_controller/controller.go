package country_controller

import (
	"backPet0/dtos/country_dto"
	"backPet0/logger"
	"backPet0/logic/country_service"
	"encoding/json"
	"net/http"
	"strconv"
)

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		logger.Log("Method not allowed: " + r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func Save(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodPost) {
		return
	}

	dto, err := country_dto.NewFromJSON(r.Body)
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = country_service.Save(dto)
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("The country was created"))
}

func GetById(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dto country_dto.CountryDTO
	dto, err = country_service.GetById(id)

	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(dto)

	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	countries, err := country_service.GetAll()
	if err != nil {
		logger.Log("Method not allowed: " + r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(countries)

	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodDelete) {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = country_service.Remove(id)

	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodPut) {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	country, err := country_dto.NewFromJSON(r.Body)
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = country_service.Update(id, country)
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
