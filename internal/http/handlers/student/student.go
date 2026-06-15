package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/irajspace/golang-crud/internal/types"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
			
		}
		slog.Info("handling request for /api/students", "method", r.Method, "url", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Student!"))
	}

}