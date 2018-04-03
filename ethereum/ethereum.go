package ethereum

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// post Пост запрос с параметрами в тело
func post(url1 string, data url.Values) string {
	form := data
	body1 := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", url1, body1)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

/*
	Все взимодействия с блокчейном Ethereum происходят через post request к nodeJS серверу
*/

// CreatePrvtKey Создание секретного ключа эфира
func CreatePrvtKey() string {
	postData := url.Values{
		"nil": {},
	}
	prvtKey := post("http://localhost:3000/createEthAccount", postData)
	return prvtKey
}

// GetAddress Получение адреса эфира по секретному ключу
func GetAddress(prvtKey string) string {
	postData := url.Values{
		"prvtKey": {prvtKey},
	}
	address := post("http://localhost:3000/getAddress", postData)
	return address
}

// GetBalance Получение баланса эфира по адресу
func GetBalance(address string) string {
	postData := url.Values{
		"address": {address},
	}
	balance := post("http://localhost:3000/getBalance", postData)
	return balance
}

// SendTransaction Отправка транзакций в блокчейн эфира
func SendTransaction(prvtKey string, sender string, receiver string, amount string) string {
	postData := url.Values{
		"prvtKey":  {prvtKey},
		"sender":   {sender},
		"receiver": {receiver},
		"amount":   {amount},
	}
	status := post("http://localhost:3000/sendTx", postData)
	println(status)
	return status
}
