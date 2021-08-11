package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidfragalaureano/sentry/service"

	"github.com/davidfragalaureano/sentry/errors"
	"github.com/gorilla/mux"
)

type sentryController struct{}

var (
	sentryService service.SentryService
)

type SentryController interface {
	Hello(w http.ResponseWriter, r *http.Request)
	GetAllSpaceObjects(w http.ResponseWriter, r *http.Request)
	GetObjectByName(w http.ResponseWriter, r *http.Request)
}

func NewSentryController(service service.SentryService) SentryController {
	sentryService = service
	return &sentryController{}
}

func (*sentryController) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sentryController{})
}

func (*sentryController) GetAllSpaceObjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	spaceObjects, err := sentryService.GetAllObjects()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"result":"","error":%q}`, err)
	}

	json.NewEncoder(w).Encode(spaceObjects)

}

func (*sentryController) GetObjectByName(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)

	if objectName, ok := pathVars["objectName"]; ok {
		w.Header().Set("Content-Type", "application/json")
		spaceObject, err := sentryService.GetObjectByName(objectName)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"result":"","error":%q}`, err)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(spaceObject)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	err := errors.NewUnexpectedError(nil, "objectName is a required variable")
	fmt.Fprintf(w, `{"result":"","error":%q}`, err)

}
