package microservices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		fmt.Fprint(w, `{error: No search criteria found, you can either search by id via /id/4 to search by name via /name/codplayconcert}`)
		return
	}
	seachkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search key found, you can either search by id via /id/4 to search by name via /name/codplayconcert}`)
		return
	}
	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
		case "name": event, err = eh.dbhandler.FindEventByName(searchkey)
		case "id": id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to find all available events %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while decoding event data %s}" )
	}
}

func newEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler {
		dbhandler: databasehandler,
	}
}
