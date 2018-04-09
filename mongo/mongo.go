package mongo

import (
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
	FoundationName   string  `bson:"foundationName"`
	Currency         string  `bson:"currency"`
	InvestInCurrency float64 `bson:"investInCurrency"`
	InvestInRub      float64 `bson:"investInRub"`
}

// Users Структура пользователя
type Users struct {
	UserID      string               `bson:"userID"`
	Name        string               `bson:"name"`
	EthPrvKey   string               `bson:"ethPrvKey"`
	EthAddress  string               `bson:"ethAddress"`
	Foundations []investInFoundation `bson:"foundations"`
}

// ConnectToMongo mongo connection
func ConnectToMongo() *mgo.Session {
	session, err := mgo.Dial("mongodb://erage:doBH8993nnjdoBH8993nnj@51.144.89.99:27017")
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

	c := session.DB("ImCup").C("foundations")
	err := c.Insert(&Foundations{Name: name, FoundedDate: foundedDate, Capital: capital, Country: country, Mission: mission})

	if err != nil {
		// log.Fatal(err)
	}
}

// AddUser Добавление пользователя
func AddUser(openSession *mgo.Session, userID string, name string, ethPrvKey string, ethAddress string) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("ImCup").C("users")

	//var foundationNullArray []investInFoundation

	err := c.Insert(&Users{UserID: userID, Name: name, EthPrvKey: ethPrvKey, EthAddress: ethAddress})

	if err != nil {
		// log.Fatal(err)
	}
}

// AddFoundationToUser Добавление благотворительной организации в БД
func AddFoundationToUser(openSession *mgo.Session, userID string, foundationName string, currency string, investInCurrency float64, investSumRub float64) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("ImCup").C("users")
	results := Users{}
	c.Find(bson.M{"UserID": userID}).One(&results)

	// Массив с фондами человека
	arr1 := results.Foundations

	// Новая структура для добавленного фонда
	arr3 := investInFoundation{FoundationName: foundationName, Currency: currency, InvestInCurrency: investInCurrency, InvestInRub: investSumRub}

	// Добавление структуры для фонда
	arr2 := append(arr1, arr3)

	err := c.Update(bson.M{"UserID": userID}, bson.M{"$set": bson.M{"foundations": arr2}})

	if err != nil {
		// log.Fatal(err)
	}
}

// FindAllFoundations Поиск всех фондов
func FindAllFoundations(openSession *mgo.Session) []Foundations {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("ImCup").C("foundations")

	var results []Foundations
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		// log.Fatal(err)
	}

	return results
}

// FindAllUsers Поиск всех users
func FindAllUsers(openSession *mgo.Session) []Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("ImCup").C("users")

	var results []Users
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		// log.Fatal(err)
	}

	return results
}

// FindUser Поиск конкретного пользователя
func FindUser(openSession *mgo.Session, userid string) Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("ImCup").C("users")

	var results Users
	err := c.Find(bson.M{"userID": userid}).One(&results)

	if err != nil {
		// log.Fatal(err)
	}

	return results
}
