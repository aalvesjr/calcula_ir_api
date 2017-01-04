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
	sal, des := getVars(r)

	s := salary.NewSalary(sal, des)
	j, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getVars(r *http.Request) (salario, descontos float32) {
	vars := mux.Vars(r)

	querySalario, ok := vars["salario"]
	if ok {
		temp, _ := strconv.ParseFloat(querySalario, 32)
		salario = float32(temp)
	}
	queryDescontos, ok := vars["descontos"]
	if ok {
		temp, _ := strconv.ParseFloat(queryDescontos, 32)
		descontos = float32(temp)
	}
	return
}
