package service

import (
	"AvoxiCodingChallenge/models"
	"encoding/json"
	"net/http"
)

func (s Service) CheckIP(w http.ResponseWriter, r *http.Request) {

	var req models.IPCheckerReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	countryMap, err := s.IPChecker(req.IP, req.Countries...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res := models.NewIPCheckerRes(countryMap)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
