package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/aalvesjr/salary"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var varsTests = []struct {
	url      string
	value    float32
	discount float32
}{
	{"/api/salario/0", 0, 0},
	{"/api/salario/1000", 1000, 0},
	{"/api/salario/1000/descontos/100", 1000, 100},
	{"/api/descontos/100", 0, 100},
}

func TestCalculateIRVars(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/salario/{salario}", CalculateIRVars).Methods("GET")
	r.HandleFunc("/api/salario/{salario}/descontos/{descontos}", CalculateIRVars).Methods("GET")
	r.HandleFunc("/api/descontos/{descontos}", CalculateIRVars).Methods("GET")

	for _, q := range varsTests {
		req, err := http.NewRequest("GET", q.url, nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
		}

		s := salary.NewSalary(q.value, q.discount)
		j, _ := json.Marshal(s)
		var b bytes.Buffer
		b.Write(j)

		if w.Body.String() != b.String() {
			t.Errorf("Salary resource expected: %v, got: %v", b.String(), w.Body.String())
		}
	}
}
