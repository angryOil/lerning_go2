package httpTest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type info struct {
	expression string
	code       int
	body       string
}

var in info

func Test_remoteSolver_Resolve(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expression := r.URL.Query().Get("expression")
			if expression != in.expression {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid expression" + in.expression))
				return
			}
			w.WriteHeader(in.code)
			w.Write([]byte(in.body))
		}),
	)
	defer server.Close()
	rs := remoteSolver{
		MathSeverURL: server.URL,
		Client:       server.Client(),
	}
	data := []struct {
		name   string
		io     info
		result float64
	}{
		{"case1", info{"2+2*10", http.StatusOK, "22"}, 22},
		{"case2", info{"2*10", http.StatusOK, "201"}, 201},
		{"case3", info{"4+2*10", http.StatusOK, "24"}, 24},
		{"case4", info{"5+2*10", http.StatusOK, "25"}, 25},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			in = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("in %f , got %f ", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != "" {
				t.Errorf("expected error \"\" but exture %s ", errMsg)
			}
		})
	}
}
