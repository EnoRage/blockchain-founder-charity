package ethereum

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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

func CreatePrvtKey() string {
	postData := url.Values{
		"nil": {},
	}
	prvtKey := post("http://localhost:3000/createEthAccount", postData)
	return prvtKey
}

func GetAddress(prvtKey string) string {
	postData := url.Values{
		"prvtKey": {prvtKey},
	}
	address := post("http://localhost:3000/getAddress", postData)
	return address
}

func GetBalance(address string) string {
	postData := url.Values{
		"address": {address},
	}
	balance := post("http://localhost:3000/getBalance", postData)
	return balance
}
