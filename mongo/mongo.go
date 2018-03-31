package mongo

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type foundations struct {
	Name        string  `bson:"Name"`
	FoundedDate int32   `bson:"FoundedDate"`
	Capital     float32 `bson:"Capital"`
	Country     string  `bson:"Country"`
	Mission     string  `bson:"Mission"`
}

type users struct {
	UserID     string `bson:"UserID"`
	Name       string `bson:"Name"`
	EthPubKey  string `bson:"EthPubKey"`
	EthPrvKey  string `bson:"EthPrvKey"`
	EthAddress string `bson:"EthAddress"`
}

// ConnectToMongo mongo connection
func ConnectToMongo() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://admin:itss2018!@medicineassistantdb.westeurope.cloudapp.azure.com:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session, err
}

// CloseMongoConnection mongo close connection
func CloseMongoConnection(session *mgo.Session) {
	session.Close()
}

// AddFoundation Добавление фонда
func AddFoundation(name string, foundedDate int32, capital float32, country string, mission string) {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")
	err = c.Insert(&foundations{Name: name, FoundedDate: foundedDate, Capital: capital, Country: country, Mission: mission})

	if err != nil {
		log.Fatal(err)
	}
}

// AddUser Добавление пользователя
func AddUser(userID string, name string, ethPubKey string, ethPrvKey string, ethAddress string) {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")
	err = c.Insert(&users{UserID: userID, Name: name, EthPubKey: ethPubKey, EthPrvKey: ethPrvKey, EthAddress: ethAddress})

	if err != nil {
		log.Fatal(err)
	}
}

// FindAllFoundations Поиск всех фондов
func FindAllFoundations() {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("Foundations")

	if err != nil {
		log.Fatal(err)
	}

	item := foundations{}

	find := c.Find(bson.M{})

	items := find.Iter()
	for items.Next(&item) {
		fmt.Println(item)
	}
}
