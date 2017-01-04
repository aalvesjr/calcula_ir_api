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

var queryTests = []struct {
	url      string
	value    float32
	discount float32
}{
	{"/api/calculate-ir", 0, 0},
	{"/api/calculate-ir?salary=1000", 1000, 0},
	{"/api/calculate-ir?salary=1000&discounts=100", 1000, 100},
	{"/api/calculate-ir?discounts=100", 0, 100},
}

func TestCalculateIRQuery(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/calculate-ir", CalculateIRQuery).Methods("GET")

	for _, q := range queryTests {
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
