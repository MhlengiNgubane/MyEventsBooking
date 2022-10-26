package microservices

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Event struct {
	ID        bson.ObjectID `bson: "_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location, omitempty"`
	Capacity int    `json:"capacity"`
}

type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

http.ListenAndServe(":8181", r)

func ServeAPI(endpoint string) error {
	handler := &eventServiceHandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.findEventHandler)

	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)

	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}

type eventServiceHandler struct{
	dbhandler persistence.DatabaseHandler
}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars = mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search key found, you can either search by id via /id/4 to search by name }`)
	}
}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {

}

func newEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler {
		dbhandler: databasehandler,
	}
}
