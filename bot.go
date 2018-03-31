package main

import (
	"log"
	"time"

	"./course"
	"github.com/tidwall/gjson"

	tb "gopkg.in/tucnak/telebot.v2"
)

var courseJSON []byte
var courseArray gjson.Result
var courseInterface interface{}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY",

		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	// Тест курсов
	// Получаем JSON формат курсов
	courseJSON = course.Course("USD")
	// Берём конкретное значение по кусам
	courseArray = gjson.Get(string(courseJSON), "WAVES")
	println(courseArray.String())

	replyBtn1 := tb.ReplyButton{Text: "💳 Мой кабинет"}
	replyBtn2 := tb.ReplyButton{Text: "Сделать пожертвование"}
	replyKeys := [][]tb.ReplyButton{
		{replyBtn1, replyBtn2},
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Главное меню", &tb.ReplyMarkup{ReplyKeyboard: replyKeys})
	})
	b.Handle(&replyBtn2, func(m *tb.Message) {
		b.Send(m.Sender, "Список благотворительных организаций: ")
	})
	b.Start()
}
