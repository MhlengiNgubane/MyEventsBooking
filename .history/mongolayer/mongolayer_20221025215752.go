package mongolayer

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB = "myevents"
	USERS = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err!= nil {
		return nil, err
	}
	return &MongoDBLayer {
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	// Let's assume the method below checks if the ID is valid for the location object of the event
	if !e.Location.ID.Valid() {
		e.LocationID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	events := []persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All()
}