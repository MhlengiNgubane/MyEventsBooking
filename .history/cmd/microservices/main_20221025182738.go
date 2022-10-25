package microservices

import (
	"net/http"

	"github.com/gorilla/mux"
)

type eventServiceHandler struct{}

type DatabaseHandler interface {
	AddEvent (Event) ([]byte, error)
	FindEvent ([]byte) (Event, error)
	Find
}

func ServeAPII(endpoint string) error {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.findEventHandler)

	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}
