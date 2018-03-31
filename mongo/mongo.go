package mongo

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Foundations struct {
	Name        string  `bson:"Name"`
	FoundedDate int32   `bson:"FoundedDate"`
	Capital     float32 `bson:"Capital"`
	Country     string  `bson:"Country"`
	Mission     string  `bson:"Mission"`
}

// mongo ferfer
func ConnectToMongo() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://admin:itss2018!@medicineassistantdb.westeurope.cloudapp.azure.com:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session, err
}

func CloseMongoConnection(session *mgo.Session) {
	session.Close()
}

func AddFoundation(name string, foundedDate int32, capital float32, country string, mission string) {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("Foundations")
	err = c.Insert(&Foundations{Name: name, FoundedDate: foundedDate, Capital: capital, Country: country, Mission: mission})

	if err != nil {
		log.Fatal(err)
	}
}

func FindAllFoundations() {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("Foundations")
	// result := []Foundations{}
	// var result []Foundations
	// c.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	item := Foundations{}

	find := c.Find(bson.M{})

	items := find.Iter()
	for items.Next(&item) {
		fmt.Println(item)
	}

	// fmt.Println(result)
}
