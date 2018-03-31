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



app.listen(3000);