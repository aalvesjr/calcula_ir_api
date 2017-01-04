package handlers

import (
	"encoding/json"
	"github.com/aalvesjr/salary"
	"net/http"
	"strconv"
)

func CalculateIRQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sal, des := getQuery(r)

	s := salary.NewSalary(sal, des)
	j, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getQuery(r *http.Request) (salario, descontos float32) {
	query := r.URL.Query()

	querySalario, ok := query["salario"]
	if ok {
		temp, _ := strconv.ParseFloat(querySalario[0], 32)
		salario = float32(temp)
	}
	queryDescontos, ok := query["descontos"]
	if ok {
		temp, _ := strconv.ParseFloat(queryDescontos[0], 32)
		descontos = float32(temp)
	}
	return
}
