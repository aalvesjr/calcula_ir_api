package handlers

import (
	"encoding/json"
	"github.com/aalvesjr/salary"
	"net/http"
	"strconv"
)

func CalculateIRQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sal, dis := getQuery(r)

	s := salary.NewSalary(sal, dis)
	j, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getQuery(r *http.Request) (salary, discounts float32) {
	query := r.URL.Query()

	querySalary, ok := query["salary"]
	if ok {
		temp, _ := strconv.ParseFloat(querySalary[0], 32)
		salary = float32(temp)
	}
	queryDiscounts, ok := query["discounts"]
	if ok {
		temp, _ := strconv.ParseFloat(queryDiscounts[0], 32)
		discounts = float32(temp)
	}
	return
}
