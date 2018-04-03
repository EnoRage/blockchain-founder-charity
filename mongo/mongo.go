package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Foundations Структура фондов
type Foundations struct {
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

// Users Структура пользователя
type Users struct {
	UserID      string               `bson:"UserID"`
	Name        string               `bson:"Name"`
	EthPrvKey   string               `bson:"EthPrvKey"`
	EthAddress  string               `bson:"EthAddress"`
	Foundations []investInFoundation `bson:"Foundations"`
}

// ConnectToMongo mongo connection
func ConnectToMongo() *mgo.Session {
	session, err := mgo.Dial("mongodb://admin:itss2018!@medicineassistantdb.westeurope.cloudapp.azure.com:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

// CloseMongoConnection mongo close connection
func CloseMongoConnection(session *mgo.Session) {
	session.Close()
}

// AddFoundation Добавление фонда
func AddFoundation(openSession *mgo.Session, name string, foundedDate int32, capital float32, country string, mission string) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")
	err := c.Insert(&Foundations{Name: name, FoundedDate: foundedDate, Capital: capital, Country: country, Mission: mission})

	if err != nil {
		log.Fatal(err)
	}
}

// AddUser Добавление пользователя
func AddUser(openSession *mgo.Session, userID string, name string, ethPrvKey string, ethAddress string) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	//var foundationNullArray []investInFoundation

	err := c.Insert(&Users{UserID: userID, Name: name, EthPrvKey: ethPrvKey, EthAddress: ethAddress})

	if err != nil {
		log.Fatal(err)
	}
}

// AddFoundationToUser Добавление благотворительной организации в БД
func AddFoundationToUser(openSession *mgo.Session, userID string, foundationName string, currency string, investInCurrency float64, investSumRub float64) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")
	results := Users{}
	c.Find(bson.M{"UserID": userID}).One(&results)

	// Массив с фондами человека
	arr1 := results.Foundations

	// Новая структура для добавленного фонда
	arr3 := investInFoundation{FoundationName: foundationName, Currency: currency, InvestInCurrency: investInCurrency, InvestInRub: investSumRub}

	// Добавление структуры для фонда
	arr2 := append(arr1, arr3)

	err := c.Update(bson.M{"UserID": userID}, bson.M{"$set": bson.M{"Foundations": arr2}})

	if err != nil {
		log.Fatal(err)
	}
}

// FindAllFoundations Поиск всех фондов
func FindAllFoundations(openSession *mgo.Session) []Foundations {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")

	var results []Foundations
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindAllUsers Поиск всех users
func FindAllUsers(openSession *mgo.Session) []Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	var results []Users
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindUser Поиск конкретного пользователя
func FindUser(openSession *mgo.Session, userid string) Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("users")

	var results Users
	err := c.Find(bson.M{"UserID": userid}).One(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}
