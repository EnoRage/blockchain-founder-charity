package objects

type Currency struct {
	name    string
	ticker  string
	assetId string
}

var Waves = Currency{name: "Waves", ticker: "WAVES", assetId: "WAVES"}
var Bitcoin = Currency{name: "Bitcoin", ticker: "BTC", assetId: "8LQW8f7P5d5PZM7GtZEBgaqRPGSzS3DfPuiXrURJ4AJS"}
var Ethereum = Currency{name: "Ethereum", ticker: "ETH", assetId: "474jTeYx2r2Va35794tCScAXWJG9hU2HcgxzMowaZUnu"}
var Litecoin = Currency{name: "Litecoin", ticker: "LTC", assetId: "HZk1mbfuJpmxU1Fs4AX5MWLVYtctsNcg6e2C6VKqK8zk"}
var ZCash = Currency{name: "ZCash", ticker: "ZEC", assetId: "BrjUWjndUanm5VsJkbUip8VRYy6LWJePtxya3FNv4TQa"}
var US_Dollar = Currency{name: "US Dollar", ticker: "USD", assetId: "Ft8X1v1LTa1ABafufpaCWyVj8KkaxUWE6xBhW6sNFJck"}
var Euro = Currency{name: "Euro", ticker: "EUR", assetId: "Gtb1WRznfchDnTh37ezoDTJ4wcoKaRsKqKjJjy7nm2zU"}

func GetTicker(e Currency) string {
	return e.name
}

func GetName(e Currency) string {
	return e.name
}

func GetAssetId(e Currency) string {
	return e.assetId
}
