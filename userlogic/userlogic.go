package userlogic

import (
	"../mongo"
)

// Auth Проверка на наличие пользователя в БД
func Auth(userID string) bool {
	len := len(mongo.FindUser(userID))

	if len != 0 {
		return true
	} else {
		return false
	}
}

// Register Регистрация в БД
func Register(userID string, userName string, ethPrvtkey string, ethAddress string) {

	mongo.AddUser(userID, userName, ethPrvtkey, ethAddress)
}
