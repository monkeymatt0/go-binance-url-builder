package go_binance_url_builder

const (
	//protocol
	SECURE_HTTP       = "https"
	SECURE_WEB_SOCKET = "wss"

	// keys
	PRODUCTION = "production"
	TEST       = "test"

	// hosts
	PRODUCTION_HOST = "api.binance.com"
	TEST_HOST       = "testnet.binance.vision"

	// path parameters
	API = "api"

	// path parameters > versions
	V1 = "v1"
	V2 = "v2"
	V3 = "v3" // As for now this one is in use

	// path parameters > misc
	ACCOUNT       = "account"
	BOOK_TICKER   = "bookTicker"
	DEPTH         = "depth"
	EXCHANGE_INFO = "exchangeInfo"
	ORDER         = "order"
	TICKER        = "ticker"
)
