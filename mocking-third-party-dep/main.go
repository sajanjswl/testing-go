package main

import "gopkg.in/mgo.v2"

type Thing struct{}


type DataAccessLayer interface {
	Insert(collectionName string, docs ...interface{}) error
  }





func Bar(data DataAccessLayer) Thing {
	doc := Thing{}

	data.Insert("foo", doc)
	return doc
}

func main() {
  mongoDAL,_ := NewMongoDAL("localhost", "database")
  
  Bar(mongoDAL)
}

// MongoDAL is an implementation of DataAccessLayer for MongoDB
type MongoDAL struct {
	session *mgo.Session
	dbName  string
}
// Insert stores documents in mongo
func (m MongoDAL) Insert(collectionName string,docs ...interface{}) error {
	return m.c(collectionName).Insert(docs)
}
// NewMongoDAL creates a MongoDAL
func NewMongoDAL(dbURI string, dbName string) (DataAccessLayer, error) {
	session, err := mgo.Dial(dbURI)
	mongo := MongoDAL{
		session: session,
		dbName:  dbName,
	}
	return mongo, err
}

// c is a helper method to get a collection from the session
func (m MongoDAL) c(collection string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collection)
}

