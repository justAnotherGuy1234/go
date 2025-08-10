package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func JsonReponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func ReadJson(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	return decoder.Decode(result)
}
