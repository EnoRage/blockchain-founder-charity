package course

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Course Эта функция получает курс в RUB или USD через get запрос
func Course(currency string) {
	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=" + currency + "&tsyms=WAVES,BTC,ETH,ZEC,LTC,USD,EUR")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
