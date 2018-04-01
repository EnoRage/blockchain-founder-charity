const Accounts = require('web3-eth-accounts'),
    Tx = require('ethereumjs-tx'),
    rp = require('request-promise'),
    safeMath = require('./safeMath.js');

var accounts = new Accounts('ws://localhost:8546');

const apiKeyToken = 'KF2VV1A88HNTQMKFBD3VYBD59SVKR5Z1QA';
const Network = 'api-rinkeby'

function createNewAccount() {
    let randomPrvtKey = accounts.create().privateKey;
    return randomPrvtKey.toString();
}

function getBalance(address, callback) {
    rp('https://' + Network + '.etherscan.io/api?module=account&action=balance&address=' + address + '&tag=latest&apikey=' + apiKeyToken, {
            json: true
        })
        .then((balance) => {
            callback(safeMath.weiToCurrency(balance.result));
        })
}

function getAddress(privateKey) {
    let account = accounts.privateKeyToAccount(privateKey);
    return account.address;
}

function sendTx(_prvtKey, _sender, _receiver, _amount) {
    let privateKey = new Buffer(_prvtKey.substring(2), 'hex');

    var tx = new Tx();

    // Узнаём цену газа
    rp('https://' + Network + '.etherscan.io/api?module=proxy&action=eth_gasPrice&apikey=' + apiKeyToken, {
            json: true
        })
        .then(
            (res) => {
                rp('https://' + Network + '.etherscan.io/api?module=proxy&action=eth_getTransactionCount&address=' + _sender + '&tag=latest&apikey=' + apiKeyToken, {
                        json: true
                    })
                    .then((txCount) => {
                        console.log(txCount)
                        tx.nonce = Number(txCount.result);
                        tx.gasPrice = Number(res.result);
                        tx.gasLimit = 400000;
                        tx.value = Number(_amount);
                        tx.from = String(_sender);
                        tx.to = String(_receiver);

                        tx.sign(privateKey);

                        const serializedTx = tx.serialize();

                        var options = {
                            uri: 'https://' + Network + '.etherscan.io/api?module=proxy&action=eth_sendRawTransaction&hex=0x' + serializedTx.toString('hex') + '&apikey=' + apiKeyToken,
                            json: true
                        }

                        rp(options)
                            .then((tx) => {
                                if (tx.error) {
                                    console.log('Транзакция не произведена');
                                    console.log(tx)
                                    return "400"


                                } else {
                                    console.log(tx);
                                    return "200"
                                }
                            })
                            .catch((err) => {
                                console.log('Транзакция не произведена')
                            });
                    });

            }
        )
}

// sendTx('0x61d94d1c3335c6c30c1336da9e4d54a586f1ffa882338a8bb9f8268296434bc9','0x6D377De54Bde59c6a4B0fa15Cb2EFB84BB32D433','0xeeD5460d1d286c406F2ADbC3CaF53ED45f898988', 20000000000000000)

module.exports.createNewAccount = createNewAccount;
module.exports.getBalance = getBalance;
module.exports.getAddress = getAddress;
module.exports.sendTx = sendTx;