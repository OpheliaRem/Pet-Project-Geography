package town_contoller

import (
	"backPet0/dtos/town_dto"
	"backPet0/logger"
	"backPet0/logic/town_service"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log("Method not allowed: " + r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	towns, err := town_service.GetAll()
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(towns)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log("Method not allowed: " + r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	town, err := town_service.GetById(id)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(town)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Log("Method not allowed: " + r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	dto := town_dto.TownDtoForCreate{}
	err := decoder.Decode(&dto)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = town_service.Save(dto)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
