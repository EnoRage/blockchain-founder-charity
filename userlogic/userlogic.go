package userlogic

import (
	"../mongo"
)

// Auth Проверка на наличие пользователя в БД
func Auth(userID string) bool {
	user := mongo.FindUser(userID)

	if user.UserID != "" {
		return true
	} else {
		return false
	}
}

// Register Регистрация в БД
func Register(userID string, userName string, ethPrvtkey string, ethAddress string) {

	mongo.AddUser(userID, userName, ethPrvtkey, ethAddress)
}
