package handlers

import (
	"encoding/json"
	"github.com/aalvesjr/salary"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CalculateIRVars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sal, dis := getVars(r)

	s := salary.NewSalary(sal, dis)
	j, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getVars(r *http.Request) (salary, discounts float32) {
	vars := mux.Vars(r)

	querySalary, ok := vars["salary"]
	if ok {
		temp, _ := strconv.ParseFloat(querySalary, 32)
		salary = float32(temp)
	}
	queryDiscounts, ok := vars["discounts"]
	if ok {
		temp, _ := strconv.ParseFloat(queryDiscounts, 32)
		discounts = float32(temp)
	}
	return
}
