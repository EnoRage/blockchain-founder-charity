const mongoose = require("mongoose"),
    Foundations = require('./schemas/foundation.js');

mongoose.Promise = global.Promise;

const user = 'admin';
const password = 'itss2018!';
const uri = 'mongodb://admin:itss2018\!@medicineassistantdb.westeurope.cloudapp.azure.com:27017/BlockChainDB?authSource=admin'; // Мейн нет ПОМЕНЯТЬ

const options = {
    autoIndex: false,
    reconnectTries: Number.MAX_VALUE,
    reconnectInterval: 500,
    poolSize: 1000,
    bufferMaxEntries: 0
};
const db = mongoose.connect(uri).then(console.log('Mongo DB works fine'));

// Foundations.find({}, (err, doc) => {
//     console.log(doc);
// });
