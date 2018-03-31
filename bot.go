package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"./course"

	tb "gopkg.in/tucnak/telebot.v2"
)

var courseJSON []byte

var courseInterface interface{}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY",

		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	courseJSON = course.Course("USD")
	// fmt.Printf("%v\n", p)
	json.Unmarshal(courseJSON, &courseInterface)
	fmt.Printf("%+v\n", courseInterface)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
