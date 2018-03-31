package main

import (
	"log"
	"time"

	"./course"
<<<<<<< HEAD
	"github.com/tidwall/gjson"
=======
	"./orglist"
>>>>>>> 6bc76a11dd3d6052192b3c11a5275fa460375f6e

	tb "gopkg.in/tucnak/telebot.v2"
)

var courseJSON []byte
var courseArray gjson.Result
var courseInterface interface{}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY",
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
	inlineBtn1 := tb.InlineButton{Unique: "0", Text: "0️⃣"}
	inlineBtn2 := tb.InlineButton{Unique: "1", Text: "1️⃣"}
	inlineBtn3 := tb.InlineButton{Unique: "2", Text: "2️⃣"}
	inlineBtn4 := tb.InlineButton{Unique: "3", Text: "3️⃣"}
	inlineBtn5 := tb.InlineButton{Unique: "4", Text: "4️⃣"}
	inlineBtn6 := tb.InlineButton{Unique: "5", Text: "5️⃣"}
	inlineBtn7 := tb.InlineButton{Unique: "6", Text: "6️⃣"}
	inlineBtn8 := tb.InlineButton{Unique: "7", Text: "7️⃣"}
	inlineBtn9 := tb.InlineButton{Unique: "8", Text: "8️⃣"}

	inlineKbrdCalc := [][]tb.InlineButton{
		{inlineBtn1, inlineBtn2, inlineBtn3, inlineBtn4, inlineBtn5, inlineBtn6, inlineBtn7, inlineBtn8, inlineBtn9},
	}
	course.Course("USD")

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
		b.Send(m.Sender, orglist.Data, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
	})
	b.Start()
}
