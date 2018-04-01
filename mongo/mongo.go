package mongo

import (
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
func AddUser(userID string, name string, ethPrvKey string, ethAddress string) {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")
	err = c.Insert(&users{UserID: userID, Name: name, EthPrvKey: ethPrvKey, EthAddress: ethAddress})

	if err != nil {
		log.Fatal(err)
	}
}

// FindAllFoundations Поиск всех фондов
func FindAllFoundations() []foundations {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")

	if err != nil {
		log.Fatal(err)
	}

	var results []foundations
	c.Find(bson.M{}).All(&results)

	return results
}

// FindAllUsers Поиск всех users
func FindAllUsers() []users {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	if err != nil {
		log.Fatal(err)
	}

	var results []users
	c.Find(bson.M{}).All(&results)

	return results
}

// FindAllUsers Поиск всех users
func FindUser(userid string) []users {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	if err != nil {
		log.Fatal(err)
	}

	var results []users
	c.Find(bson.M{"UserID": userid}).All(&results)

	return results
}
