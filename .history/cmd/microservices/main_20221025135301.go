package microservices

import(
	"github.com/gorilla/mux"
)

r := mux.NewRouter()
eventsrouter := r.PathPrefix("/events").Subrouter()

eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").
HandlerFunc(handler.findEventHandler)

eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

eventsrouter.Methods("POST")