package microservices

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ServeAPII(endpoint string) error {
	handler := &eventServiceHandler{}
r := mux.NewRouter()
eventsrouter := r.PathPrefix("/events").Subrouter()

eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").
HandlerFunc(handler.findEventHandler)

eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
}
type eventServiceHandler struct {}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}

