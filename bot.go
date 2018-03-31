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
var sum = ""

// для диалогов переменные

var courseJSON []byte
var courseResult gjson.Result

var wavesBalanceResult gjson.Result
var currencyBalanceResult gjson.Result

func main() {
	b, err := tb.NewBot(tb.Settings{
		// Token: "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY",
		Token:  "525513661:AAEdYAbizNP8SiT2fhjweHRZULFL84KsUYk",
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

	//Тест mongoDB
	// Добавление фонда
	// mongo.AddFoundation("Имя", 2018, 1.3, "Россия", "Информация о фонде")
	// Поиск по фондам
	// mongo.FindAllFoundations()

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
		[]tb.ReplyButton{replyBtn1},
		[]tb.ReplyButton{replyBtn2},
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

	inlineklav0 := tb.InlineButton{Unique: "klav0", Text: "0️⃣"}
	inlineklav1 := tb.InlineButton{Unique: "klav1", Text: "1️⃣"}
	inlineklav2 := tb.InlineButton{Unique: "klav2", Text: "2️⃣"}
	inlineklav3 := tb.InlineButton{Unique: "klav3", Text: "3️⃣"}
	inlineklav4 := tb.InlineButton{Unique: "klav4", Text: "4️⃣"}
	inlineklav5 := tb.InlineButton{Unique: "klav5", Text: "5️⃣"}
	inlineklav6 := tb.InlineButton{Unique: "klav6", Text: "6️⃣"}
	inlineklav7 := tb.InlineButton{Unique: "klav7", Text: "7️⃣"}
	inlineklav8 := tb.InlineButton{Unique: "klav8", Text: "8️⃣"}
	inlineklav9 := tb.InlineButton{Unique: "klav9", Text: "9️⃣"}
	inlineklavdot := tb.InlineButton{Unique: "klavdot", Text: " . "}
	inlineklavapply := tb.InlineButton{Unique: "enter", Text: "✅ Подтвердить"}
	inlineklavrenew := tb.InlineButton{Unique: "renew", Text: "🆕 Заново"}
	inlineklavback := tb.InlineButton{Unique: "remove", Text: "❌ Назад"}
	inlineklavdellast := tb.InlineButton{Unique: "last", Text: "⬅️"}
	inlineKbrdsum := [][]tb.InlineButton{
		{inlineklav1, inlineklav2, inlineklav3}, {inlineklav4, inlineklav5, inlineklav6},
		{inlineklav7, inlineklav8, inlineklav9}, {inlineklavdot, inlineklav0, inlineklavdellast},
		{inlineklavrenew, inlineklavback}, {inlineklavapply},
	}

	inlinуvapply := tb.InlineButton{Unique: "apply", Text: "✅ Подтвердить"}
	inlineKbrdaply := [][]tb.InlineButton{{inlinуvapply}}
	course.Course("USD")

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{DisableWebPagePreview: true}, &tb.ReplyMarkup{ReplyKeyboard: replyKeys})
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
		b.Send(m.Sender, orglist.DataAdd, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond1", func(m *tb.Message) {
		fond = "2"
		b.Send(m.Sender, orglist.DataAdd1, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond2", func(m *tb.Message) {
		fond = "3"
		b.Send(m.Sender, orglist.DataAdd2, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond3", func(m *tb.Message) {
		fond = "4"
		b.Send(m.Sender, orglist.DataAdd3, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond4", func(m *tb.Message) {
		fond = "5"
		b.Send(m.Sender, orglist.DataAdd4, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond5", func(m *tb.Message) {
		fond = "6"
		b.Send(m.Sender, orglist.DataAdd5, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond6", func(m *tb.Message) {
		fond = "7"
		b.Send(m.Sender, orglist.DataAdd6, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond7", func(m *tb.Message) {
		fond = "8"
		b.Send(m.Sender, orglist.DataAdd7, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond8", func(m *tb.Message) {
		fond = "9"
		b.Send(m.Sender, orglist.DataAdd8, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	// слушает какой фонд выбрал

	// при нажатии кнопки пожертвовать происходит оплата
	b.Handle(&inlineInv, func(c *tb.Callback) {
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// при нажатии кнопки пожертвовать происходит оплата

	// тут клавиатурка по занесению денег
	b.Handle(&inlineklav1, func(c *tb.Callback) {
		sum += "1"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav2, func(c *tb.Callback) {
		sum += "2"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav3, func(c *tb.Callback) {
		sum += "3"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav4, func(c *tb.Callback) {
		sum += "4"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav5, func(c *tb.Callback) {
		sum += "5"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav6, func(c *tb.Callback) {
		sum += "6"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav7, func(c *tb.Callback) {
		sum += "7"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav8, func(c *tb.Callback) {
		sum += "8"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav9, func(c *tb.Callback) {
		sum += "9"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavdot, func(c *tb.Callback) {
		sum += "."
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavrenew, func(c *tb.Callback) {
		sum = ""
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavdellast, func(c *tb.Callback) {
		sz := len(sum)
		if sz > 0 {
			sum = sum[:sz-1]
		}
		var msg = orglist.EnterSum + "Текущая сумма: " + sum
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavapply, func(c *tb.Callback) {
		var msg = "Текущая сумма: " + sum + "\n\nВы готовы перевести фонду " + fond
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdaply})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// тут клавиатурка по занесению денег

	// final apply
	b.Handle(&inlinуvapply, func(c *tb.Callback) {
		var msg = "Перевод совершен успешно, подробности в личном кабинете"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{DisableWebPagePreview: true}, &tb.ReplyMarkup{ReplyKeyboard: replyKeys})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// final apply

	b.Start()
}
