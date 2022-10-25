package microservices

import(
	"github.com/gorilla/mux"
)

r := mux.NewRouter()
eventsrouter := r.PathPrefix("/events").Subrouter()

eventsrouter.Methods("GET").Path("/{SearchCriteria}/{}")