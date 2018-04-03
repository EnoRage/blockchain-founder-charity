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

type investInFoundation struct {
	FoundationName   string  `bson:"FoundationName"`
	Currency         string  `bson:"Currency"`
	InvestInCurrency float64 `bson:"InvestInCurrency"`
	InvestInRub      float64 `bson:"InvestInRub"`
}

type users struct {
	UserID      string               `bson:"UserID"`
	Name        string               `bson:"Name"`
	EthPrvKey   string               `bson:"EthPrvKey"`
	EthAddress  string               `bson:"EthAddress"`
	Foundations []investInFoundation `bson:"Foundations"`
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

	c := session.DB("BlockChainDB").C("users")

	//var foundationNullArray []investInFoundation

	err = c.Insert(&users{UserID: userID, Name: name, EthPrvKey: ethPrvKey, EthAddress: ethAddress})

	if err != nil {
		log.Fatal(err)
	}
}

// AddFoundationToUser Добавление благотворительной организации в БД
func AddFoundationToUser(userID string, foundationName string, currency string, investInCurrency float64, investSumRub float64) {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")
	results := users{}
	c.Find(bson.M{"UserID": userID}).One(&results)

	// Массив с фондами человека
	arr1 := results.Foundations

	// Новая структура для добавленного фонда
	arr3 := investInFoundation{FoundationName: foundationName, Currency: currency, InvestInCurrency: investInCurrency, InvestInRub: investSumRub}

	// Добавление структуры для фонда
	arr2 := append(arr1, arr3)

	err = c.Update(bson.M{"UserID": userID}, bson.M{"$set": bson.M{"Foundations": arr2}})

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

// FindUser Поиск конкретного пользователя
func FindUser(userid string) users {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	if err != nil {
		log.Fatal(err)
	}

	var results users
	c.Find(bson.M{"UserID": userid}).One(&results)

	return results
}
