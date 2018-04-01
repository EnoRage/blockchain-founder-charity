package mongo

import (
	"fmt"
	"log"
	"strconv"

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
	UserID      string     `bson:"UserID"`
	Name        string     `bson:"Name"`
	EthPrvKey   string     `bson:"EthPrvKey"`
	EthAddress  string     `bson:"EthAddress"`
	Foundations [][]string `bson:"Foundations"`
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

	var foundationNullArray [][]string

	err = c.Insert(&users{UserID: userID, Name: name, EthPrvKey: ethPrvKey, EthAddress: ethAddress, Foundations: foundationNullArray})

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

	var counter int = 0
	for i := 0; i < len(results.Foundations); i++ {
		if len(results.Foundations) != 0 {
			if results.Foundations[i][0] != "" {
				counter++
			} else {
				break
			}
		} else {
			break
		}

	}
	println(string(counter))

	arr1 := results.Foundations
	// var arr2 [][]string
	arr3 := []string{foundationName, currency, strconv.FormatFloat(investInCurrency, 'g', 8, 64), strconv.FormatFloat(investSumRub, 'g', 8, 64)}
	// copy(arr2, arr1)
	arr2 := append(arr1, arr3)
	// arr2[counter] = arr3
	// fmt.Println(arr2)
	fmt.Println(arr2)
	// arr2[counter][1] =
	// arr2[counter][2] =
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

// FindUserFoundations Поиск всех фондов юзера
func FindUserFoundations(userid string) []foundations {
	session, err := ConnectToMongo()
	defer CloseMongoConnection(session)

	c := session.DB("BlockChainDB").C("foundations")

	if err != nil {
		log.Fatal(err)
	}

	var results []foundations
	c.Find(bson.M{"UserID": userid}).All(&results)

	return results
}
