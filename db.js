const mongoose = require("mongoose"),
    Foundations = require('./schemas/foundation.js');

mongoose.Promise = global.Promise;

<<<<<<< HEAD
const uri = 'mongodb://erage:doBH8993nnjdoBH8993nnj@51.144.89.99:27017/ImCup?authSource=admin'; // Мейн нет ПОМЕНЯТЬ

=======
// const user = 'admin';
// const password = 'itss2018!';
// const uri = 'mongodb://admin:itss2018\!@medicineassistantdb.westeurope.cloudapp.azure.com:27017/BlockChainDB?authSource=admin'; // Мейн нет ПОМЕНЯТЬ
// const uri = 'mongodb://51.144.89.99:27017/Charity'; // Мейн нет ПОМЕНЯТЬ
>>>>>>> 090e816fd1f105d4fce912b40fbe98f29c3cd3d4
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
