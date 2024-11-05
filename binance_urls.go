package go_binance_url_builder

const (
	//protocol
	SECURE_HTTP       = "https"
	SECURE_WEB_SOCKET = "wss"

	// keys
	PRODUCTION     = "production"
	PRODUCTION_WSS = "production_wss"
	TEST           = "test"
	TEST_WSS       = "test_wss"

	// hosts
	PRODUCTION_HOST    = "api.binance.com"
	PRODUCTION_WS_HOST = "ws-api.binance.com"
	TEST_HOST          = "testnet.binance.vision"
	TEST_WS_HOST       = "ws-api.testnet.binance.vision"

	PRODUCTION_WSS_HOST = "stream.binance.com:443"

	// path parameters
	API     = "api"
	WSS_API = "ws-api"

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
