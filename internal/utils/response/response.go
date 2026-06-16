package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)
type Response struct {	
	Status    string         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}	
const (
	StatusOK="OK"
	StatusError="ERROR"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error)Response{
	return Response{
		Status:   StatusError,
		Error: err.Error(),
		
	}
}
func ValidationError(err validator.ValidationErrors)Response{
	var errorMessages []string
	for _, fieldErr := range err {
		switch fieldErr.Tag() {
		case "required":
			errorMessages = append(errorMessages, fieldErr.Field()+" is required")
		default:
			errorMessages = append(errorMessages, fieldErr.Field()+" is invalid")
		}
	}
	return Response{
		Status:   StatusError,
		Error: "Validation failed: " +  strings.Join(errorMessages, ", "),
		
	}
	
}