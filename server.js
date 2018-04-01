const express = require('express'),
    url = require("url"),
    bodyParser = require('body-parser'),
    Ethereum = require('./ethereum.js'),
    db = require('./db.js');


const app = express();

// Чтоб считывать статические файлы
app.use(bodyParser.urlencoded({
    extended: false
}))

// Создаём секретный ключ эфира
app.post('/createEthAccount', (req, res) => {
    var account = Ethereum.createNewAccount();
    res.send(account);
});

// Получаем адрес из секретного ключа
app.post('/getAddress', (req, res) => {
    var data = req.body;
    console.log(data)
    var address = Ethereum.getAddress(data.prvtKey);
    res.send(address);
});

// Получаем баланс эфира
app.post('/getBalance', (req, res) => {
    var data = req.body;
    console.log(data)
    var address = Ethereum.getBalance(data.address, (balance) => {
        console.log(balance)
        res.send(balance.toString());
    });
});

// Отправляем транзакцию в блокчейн
app.post('/sendTx', (req, res) => {
    var data = req.body;
    console.log(data);
    Ethereum.sendTx(data.prvtKey, data.sender, data.receiver, data.amount);
    res.send('200');
});

app.listen(3000);