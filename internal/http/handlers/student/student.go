package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/irajspace/golang-crud/internal/types"
	"github.com/irajspace/golang-crud/internal/utils/response"
)
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
 
		if errors.Is(err, io.EOF) {
			response.WriteJSON(
				w,
				http.StatusBadRequest,
				response.GeneralError(fmt.Errorf("empty request body")),
			)
			return
		}

		if err != nil {
			response.WriteJSON(
				w,
				http.StatusBadRequest,
				response.GeneralError(errors.New("invalid JSON")),
			)
		
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validationErrs := err.(validator.ValidationErrors)
			response.WriteJSON(
				w,
				http.StatusBadRequest,
				response.ValidationError(validationErrs),
			)
			return
		}

		slog.Info(
			"student created",
			"method", r.Method,
			"url", r.URL.Path,
			"student", student,
		)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Student!"))
	}
}