package main

import (
	"log"
	"time"

	"./course"
	"./objects"
	"./orglist"
	"./waves"
	"github.com/tidwall/gjson"

	tb "gopkg.in/tucnak/telebot.v2"
)

// для диалогов переменные
var fond = ""

// для диалогов переменные

var courseJSON []byte
var courseResult gjson.Result

var wavesBalanceResult gjson.Result
var currencyBalanceResult gjson.Result

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY",

		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	// Тест объектов
	// Получаем assetId конкретной валюты
	assetId := objects.GetAssetId(objects.Bitcoin)
	// Получаем name конкретной валюты
	name := objects.GetName(objects.Bitcoin)
	// Получаем ticker конкретной валюты
	ticker := objects.GetTicker(objects.Bitcoin)

	// Тест курсов
	// Получаем JSON формат курсов
	courseJSON = course.Course("USD")
	/*
		Берём конкретное значение по кусам.
		Доступные параметры: WAVES, BTC, ETH, ZEC, LTC, USD, EUR
	*/
	courseResult = gjson.Get(string(courseJSON), "WAVES")

	// Тест Waves
	// Получаем баланс аккаунта в WAVES
	wavesBalanceResult = waves.GetWavesBalance("3P3Xd5x7hqKjhKhJXKfPo1RVhixTCWA9TE2")
	// Получаем баланс аккаунта в другой валюте
	currencyBalanceResult = waves.GetBalance("3P3Xd5x7hqKjhKhJXKfPo1RVhixTCWA9TE2", objects.GetAssetId(objects.ZCash))

	// Тестовые логи
	println(assetId)
	println(name)
	println(ticker)
	println(courseResult.String())
	println(wavesBalanceResult.String())
	println(currencyBalanceResult.String())

	replyBtn1 := tb.ReplyButton{Text: "💳 Мой кабинет"}
	replyBtn2 := tb.ReplyButton{Text: "Сделать пожертвование"}
	replyKeys := [][]tb.ReplyButton{
		{replyBtn1, replyBtn2},
	}
	inlineBtn1 := tb.InlineButton{Unique: "1", Text: "1️⃣"}
	inlineBtn2 := tb.InlineButton{Unique: "2", Text: "2️⃣"}
	inlineBtn3 := tb.InlineButton{Unique: "3", Text: "3️⃣"}
	inlineBtn4 := tb.InlineButton{Unique: "4", Text: "4️⃣"}
	inlineBtn5 := tb.InlineButton{Unique: "5", Text: "5️⃣"}
	inlineBtn6 := tb.InlineButton{Unique: "6", Text: "6️⃣"}
	inlineBtn7 := tb.InlineButton{Unique: "7", Text: "7️⃣"}
	inlineBtn8 := tb.InlineButton{Unique: "8", Text: "8️⃣"}
	inlineBtn9 := tb.InlineButton{Unique: "9", Text: "9️⃣"}

	inlineKbrdCalc := [][]tb.InlineButton{
		{inlineBtn1, inlineBtn2, inlineBtn3, inlineBtn4, inlineBtn5, inlineBtn6, inlineBtn7, inlineBtn8, inlineBtn9},
	}

	inlineInv := tb.InlineButton{Unique: "inv", Text: "Перевести"}
	inlineInvMenu := [][]tb.InlineButton{
		{inlineInv},
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
	// тут переход в список фондов с пожертвованиями
	b.Handle(&replyBtn2, func(m *tb.Message) {
		b.Send(m.Sender, orglist.Data, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
	})
	// тут переход в список фондов с пожертвованиями

	// inline buttons 1-9 Инфа о фондах
	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data1, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data2, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn3, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data3, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn4, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data4, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn5, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data5, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn6, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data6, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn7, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data7, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn8, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data8, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineBtn9, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data9, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// inline buttons 1-9

	// слушает какой фонд выбрал
	b.Handle("/fond0", func(m *tb.Message) {
		fond = "1"
		b.Send(m.Sender, "0 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond1", func(m *tb.Message) {
		fond = "2"
		b.Send(m.Sender, "Первый фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond2", func(m *tb.Message) {
		fond = "3"
		b.Send(m.Sender, "2 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond3", func(m *tb.Message) {
		fond = "4"
		b.Send(m.Sender, "3 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond4", func(m *tb.Message) {
		fond = "5"
		b.Send(m.Sender, "4 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond5", func(m *tb.Message) {
		fond = "6"
		b.Send(m.Sender, "5 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond6", func(m *tb.Message) {
		fond = "7"
		b.Send(m.Sender, "6 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond7", func(m *tb.Message) {
		fond = "8"
		b.Send(m.Sender, "7 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond8", func(m *tb.Message) {
		fond = "9"
		b.Send(m.Sender, "8 фонд тут кнопочки купить и подробнее инфа", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	// слушает какой фонд выбрал

	// при нажатии кнопки пожертвовать происходит оплата
	b.Handle(&inlineInv, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data9, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// при нажатии кнопки пожертвовать происходит оплата
	b.Start()
}
