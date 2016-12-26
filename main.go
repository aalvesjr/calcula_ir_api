package main

import (
	"bytes"
	"github.com/aalvesjr/calcula_ir_api/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	port = "PORT"
)

func main() {
	viper.AutomaticEnv()
	var bufferPort bytes.Buffer
	bufferPort.WriteString(":")
	bufferPort.WriteString(viper.GetString(port))

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/calcula-ir", handlers.CalculaIRQuery).Methods("GET")
	r.HandleFunc("/api/salario/{salario}", handlers.CalculaIRVars).Methods("GET")
	r.HandleFunc("/api/salario/{salario}/descontos/{descontos}", handlers.CalculaIRVars).Methods("GET")

	server := &http.Server{
		Addr:    bufferPort.String(),
		Handler: r,
	}
	log.Println("Listening on port", viper.GetString(port))
	server.ListenAndServe()
}
