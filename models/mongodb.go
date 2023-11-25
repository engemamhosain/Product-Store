// models/mongodb.go
package models

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


type Counter struct {
	ID   string `bson:"_id"`
	Seq  int    `bson:"seq"`
}


var session *mgo.Session

func init() {
    // Initialize MongoDB session
    var err error
    session, err = mgo.Dial("mongodb://dbOwner:go%40test@103.199.168.131:19928/?authSource=product_store")
    if err != nil {
        panic(err)
    }

    // Optional: Set session mode and other configurations if needed
    // session.SetMode(mgo.Monotonic, true)
    // ...
}

// GetSession returns the MongoDB session
func GetSession() *mgo.Session {
    return session.Copy()
}



func getNextSequence(session *mgo.Session, counterName string) (int) {
	collection := session.DB("product_store").C("counters")

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}

	var counter Counter
	_, err := collection.FindId(counterName).Apply(change, &counter)
	if err != nil {
		return 0
	}

	return counter.Seq
}
