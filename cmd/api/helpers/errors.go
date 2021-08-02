package helpers

import (
	"fmt"
	"log"
	"net/http"
)

type Errors struct {
	helpers Helpers
	logger  *log.Logger
}

func (errors *Errors) logError(r *http.Request, err error) {
	errors.logger.Println(err)
}

func (errors *Errors) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	env := envelope{"error": message}
	err := errors.helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		errors.logError(r, err)
		w.WriteHeader(status)
	}
}

// Send internal server error response
func (errors *Errors) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errors.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	errors.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

// Send not found error response
func (errors *Errors) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errors.ErrorResponse(w, r, http.StatusNotFound, message)
}

// Send method not allowed error response
func (errors *Errors) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errors.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// Send bad request error response
func (errors *Errors) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errors.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}
