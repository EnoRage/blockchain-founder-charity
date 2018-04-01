package main

import (
	"log"
	"math"
	"strconv"
	"time"

	"./course"
	"./ethereum"
	"./mongo"
	"./objects"
	"./orglist"
	"./userlogic"
	"./waves"
	"github.com/tidwall/gjson"
	tb "gopkg.in/tucnak/telebot.v2"
)

// для диалогов переменные
var fond = ""
var sum = ""
var concurrency = ""
var rubsum = ""

var courseJSON []byte
var courseResult gjson.Result

var wavesBalanceResult gjson.Result
var currencyBalanceResult gjson.Result

func main() {
	b, err := tb.NewBot(tb.Settings{
		// Token: "576497547:AAFqeiPb5j5fVktRPqtzpTvaIp8ExKlZZAY", //продакшн @bf_charity_bot
		Token: "525513661:AAEdYAbizNP8SiT2fhjweHRZULFL84KsUYk", //Никита @botGoTestBot.
		// Token:  "539909670:AAFk7Lxz73lTbtfjf8xIReCwSoEZZpjAlqI", //Кирилл @kirillBotGo_bot
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	// Тест объектов
	// Получаем assetId конкретной валюты
	assetID := objects.GetAssetID(objects.Bitcoin)
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
	currencyBalanceResult = waves.GetBalance("3P3Xd5x7hqKjhKhJXKfPo1RVhixTCWA9TE2", objects.GetAssetID(objects.ZCash))
	// Создание Seed
	seed := waves.CreateSeed()

	// Тест подключения nodejs
	// Получаем privateKey
	prvtKey := ethereum.CreatePrvtKey()
	// Получаем адрес
	address := ethereum.GetAddress(prvtKey)
	// Получаем баланс эфира
	balance := ethereum.GetBalance("0x6D377De54Bde59c6a4B0fa15Cb2EFB84BB32D433")
	// Отправляем транзакцию
	// status := ethereum.SendTransaction("0x61d94d1c3335c6c30c1336da9e4d54a586f1ffa882338a8bb9f8268296434bc9", "0x6D377De54Bde59c6a4B0fa15Cb2EFB84BB32D433", "0x03b825db4af2A61eaFdeCe3A2AA3039743996df2", "1000")

	//Тест mongoDB
	// Добавление фонда
	// mongo.AddFoundation("Имя", 2018, 1.3, "Россия", "Информация о фонде")
	// Поиск по фондам
	foundationCollection := mongo.FindAllFoundations()

	// Тестовые логи
	println(assetID)
	println(name)
	println(ticker)
	println(courseResult.String())
	println(wavesBalanceResult.String())
	println(currencyBalanceResult.String())
	println(seed.String())
	println(prvtKey)
	println(address)
	println(balance)
	// println(status)

	// Поиск всех имён из коллекции foundations
	for k := range foundationCollection {
		println(foundationCollection[k].Name)
	}

	replyBtn1 := tb.ReplyButton{Text: "💳 Мой кабинет"}
	replyBtn2 := tb.ReplyButton{Text: "💸 Список благотворительных организаций"}
	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{replyBtn1},
		[]tb.ReplyButton{replyBtn2},
	}
	inlineBtn0 := tb.InlineButton{Unique: "0", Text: "0️⃣"}
	inlineBtn1 := tb.InlineButton{Unique: "1", Text: "1️⃣"}
	inlineBtn2 := tb.InlineButton{Unique: "2", Text: "2️⃣"}
	inlineBtn3 := tb.InlineButton{Unique: "3", Text: "3️⃣"}
	inlineBtn4 := tb.InlineButton{Unique: "4", Text: "4️⃣"}
	inlineBtn5 := tb.InlineButton{Unique: "5", Text: "5️⃣"}
	inlineBtn6 := tb.InlineButton{Unique: "6", Text: "6️⃣"}
	// inlineBtn7 := tb.InlineButton{Unique: "7", Text: "7️⃣"}
	// inlineBtn8 := tb.InlineButton{Unique: "8", Text: "8️⃣"}

	inlineKbrdCalc := [][]tb.InlineButton{
		{inlineBtn0, inlineBtn1, inlineBtn2, inlineBtn3, inlineBtn4, inlineBtn5, inlineBtn6},
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

	inlineBtnWAV := tb.InlineButton{Unique: "WAVES", Text: "📈 WAVES"}
	inlineBtnBTC := tb.InlineButton{Unique: "BTC", Text: "📈 BTC"}
	inlineBtnETH := tb.InlineButton{Unique: "ETH", Text: "📈 ETH"}
	inlineBtnLTC := tb.InlineButton{Unique: "LTC", Text: "📈 LTC"}
	inlineCurrency := [][]tb.InlineButton{{inlineBtnWAV, inlineBtnBTC}, {inlineBtnETH, inlineBtnLTC}}
	course.Course("USD")

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		// logs
		var userID = strconv.Itoa(m.Sender.ID)
		var name = string(m.Sender.Username)
		var prvtKeyETH = ethereum.CreatePrvtKey()
		var addressETH = ethereum.GetAddress(prvtKeyETH)
		// println()
		// println(name)
		// println(prvtKeyETH)
		// println(addressETH)
		// println(userlogic.Auth(userID))
		if userlogic.Auth(userID) != true {
			userlogic.Register(userID, name, prvtKeyETH, addressETH)
			var msg = "Вы зарегистрированы в системе!\n\n"
			msg += "Ваш *Private Key:* "
			msg += prvtKeyETH
			msg += "\n\n"
			msg += "Ваш *Address:* "
			msg += addressETH
			b.Send(m.Sender, msg, &tb.SendOptions{DisableWebPagePreview: true, ParseMode: "Markdown"})
		}
		// len := len(mongo.FindUser(userID))
		// if len != 0 {
		// } else {
		// 	mongo.AddUser(userID, name, prvtKeyETH, addressETH)
		// }
		// println(len)
		// println("))))")
		// println(m.Sender.ID)
		// logs
		var text = "*Главное меню*\n\nB *Charity* - стандарт участия в благотворительности.\n\n"
		text += "*1.* Выбираешь благотворительную организацию\n"
		text += "*2.* Инвестируешь в конкретный  реальный проект\n"
		text += "*3.* Принимаешь участие в контроле проекта\n"
		text += "\n*Приемущества*\n\n"
		text += "*1.* *Прямой* и *безопасный* перевод в организацию без посредников\n"
		text += "*2.* *Контроль расходов* благотворительной организации\n"
		text += "*3.* Возможность *активного участия* и социальная ответственность\n"
		b.Send(m.Sender, text, &tb.SendOptions{DisableWebPagePreview: true, ParseMode: "Markdown"}, &tb.ReplyMarkup{ReplyKeyboard: replyKeys})
	})
	// тут переход в список фондов с пожертвованиями
	b.Handle(&replyBtn2, func(m *tb.Message) {
		b.Send(m.Sender, orglist.Data, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
	})
	b.Handle(&replyBtn1, func(m *tb.Message) {
		user := mongo.FindUser(strconv.Itoa(m.Sender.ID))
		var eth = ethereum.GetBalance(user[0].EthAddress)
		//0x7fb5f775c04b42bdc7506404272a3845d6d2e6c0be1671b24bc242f9ea43912a
		println(ethereum.GetBalance("0x7fb5f775c04b42bdc7506404272a3845d6d2e6c0be1671b24bc242f9ea43912a"))
		ethufufuuufuuf, err := strconv.ParseFloat(eth, 64)
		if err != nil {
			println(eth)
		}
		var ethreal = ethufufuuufuuf * math.Pow(10, -18)
		var thefuckingrealeth = strconv.FormatFloat(ethreal, 'g', 1, 64)
		var torub = course.Course("RUB")
		var torub3 = (1.0 / (gjson.Get(string(torub), "ETH").Float())) * ethreal
		var ethrub = strconv.FormatFloat(torub3, 'g', 1, 64)
		var msg = "*Личный кабинет* \n\n*Баланс по валютам:*" + "\n\n`ETH:` " + thefuckingrealeth + " (" + ethrub + " RUB)"
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"})
	})

	// тут переход в список фондов с пожертвованиями

	// inline buttons 1-9 Инфа о фондах

	b.Handle(&inlineBtn0, func(c *tb.Callback) {
		b.Edit(c.Message, orglist.Data, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
		b.Respond(c, &tb.CallbackResponse{})
	})
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
	// b.Handle(&inlineBtn7, func(c *tb.Callback) {
	// 	b.Edit(c.Message, orglist.Data7, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
	// 	b.Respond(c, &tb.CallbackResponse{})
	// })
	// b.Handle(&inlineBtn8, func(c *tb.Callback) {
	// 	b.Edit(c.Message, orglist.Data8, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdCalc})
	// 	b.Respond(c, &tb.CallbackResponse{})
	// })

	// inline buttons 1-6

	// слушает какой фонд выбрал
	b.Handle("/fond0", func(m *tb.Message) {
		fond = "Bill & Melinda Gates Foundation"
		b.Send(m.Sender, orglist.DataAdd, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond1", func(m *tb.Message) {
		fond = "Подари Жизнь"
		b.Send(m.Sender, orglist.DataAdd1, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond2", func(m *tb.Message) {
		fond = "Welcome Trust"
		b.Send(m.Sender, orglist.DataAdd2, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond3", func(m *tb.Message) {
		fond = "Ford Foundation"
		b.Send(m.Sender, orglist.DataAdd3, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond4", func(m *tb.Message) {
		fond = "Linux Foundation"
		b.Send(m.Sender, orglist.DataAdd4, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond5", func(m *tb.Message) {
		fond = "Ethereum Foundation"
		b.Send(m.Sender, orglist.DataAdd5, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})
	b.Handle("/fond6", func(m *tb.Message) {
		fond = "РусФонда"
		b.Send(m.Sender, orglist.DataAdd6, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineInvMenu})
	})

	// слушает какой фонд выбрал

	// тут клавиатурка по занесению денег
	b.Handle(&inlineklav0, func(c *tb.Callback) {
		sum += "0"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav1, func(c *tb.Callback) {
		sum += "1"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav2, func(c *tb.Callback) {
		sum += "2"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav3, func(c *tb.Callback) {
		sum += "3"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav4, func(c *tb.Callback) {
		sum += "4"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav5, func(c *tb.Callback) {
		sum += "5"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav6, func(c *tb.Callback) {
		sum += "6"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav7, func(c *tb.Callback) {
		sum += "7"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav8, func(c *tb.Callback) {
		sum += "8"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklav9, func(c *tb.Callback) {
		sum += "9"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavdot, func(c *tb.Callback) {
		sum += "."
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavrenew, func(c *tb.Callback) {
		sum = ""
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavdellast, func(c *tb.Callback) {
		sz := len(sum)
		if sz > 0 {
			sum = sum[:sz-1]
		}
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavapply, func(c *tb.Callback) {
		var msg = "*Данные о переводе*\n\n" + "`Организация: ` *" + fond + "*\n\n`Сумма пожертвования:` *" + sum + "*` " + concurrency + "` или *" + rubsum + "* `RUB`"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdaply})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&inlineklavback, func(c *tb.Callback) {
		sum = ""
		concurrency = ""
		var msg = "Выберите валюту для перевода: \n\n`Только для ETH доступна операция отслеживания того, что делает организация`"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineCurrency})
		b.Respond(c, &tb.CallbackResponse{})
	})

	// тут клавиатурка по занесению денег

	// Выбрать валюту после фонда
	b.Handle(&inlineInv, func(c *tb.Callback) {
		var msg = "Выберите валюту для перевода: \n\n`Только для ETH доступна операция отслеживания того, что делает организация`"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineCurrency})
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&inlineBtnWAV, func(c *tb.Callback) {
		concurrency = "Waves"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&inlineBtnBTC, func(c *tb.Callback) {
		concurrency = "Bitcoin"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&inlineBtnETH, func(c *tb.Callback) {
		concurrency = "Ethereum"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&inlineBtnLTC, func(c *tb.Callback) {
		concurrency = "Litecoin"
		var msg = orglist.EnterSum + "Текущая сумма: " + sum + " " + concurrency
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: inlineKbrdsum})
		b.Respond(c, &tb.CallbackResponse{})
	})

	// Выбрать валюту после фонда

	// final apply
	b.Handle(&inlinуvapply, func(c *tb.Callback) {
		var msg = "Перевод совершен успешно, подробности в личном кабинете"
		concurrency = ""
		sum = ""
		fond = ""
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{DisableWebPagePreview: true}, &tb.ReplyMarkup{ReplyKeyboard: replyKeys})
		b.Respond(c, &tb.CallbackResponse{})
	})
	// final apply

	b.Start()
}
