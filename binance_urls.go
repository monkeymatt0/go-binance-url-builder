package go_binance_url_builder

const (
	//protocol
	SECURE_HTTP       = "https"
	SECURE_WEB_SOCKET = "wss"

	// keys
	PRODUCTION     = "production"
	PRODUCTION_WS  = "production_ws"
	PRODUCTION_WSS = "production_wss"
	TEST           = "test"
	TEST_WS        = "test_ws"
	TEST_WSS       = "test_wss"

	// hosts
	PRODUCTION_HOST    = "api.binance.com"
	PRODUCTION_WS_HOST = "ws-api.binance.com"
	TEST_HOST          = "testnet.binance.vision"
	TEST_WS_HOST       = "ws-api.testnet.binance.vision"

	PRODUCTION_WSS_HOST = "stream.binance.com:443"             // websocket stream to retrieve infos
	TEST_WSS_HOST       = "stream.testnet.binance.vision:9443" // websocket stream to retrieve infos (TEST)

	// path parameters
	API     = "api"
	WS_API  = "ws-api"
	WSS_API = "ws"

	// path parameters > versions
	V1 = "v1"
	V2 = "v2"
	V3 = "v3" // As for now this one is in use

	// path parameters > misc
	BOOK_TICKER   = "bookTicker"
	DEPTH         = "depth"
	EXCHANGE_INFO = "exchangeInfo"
	TICKER        = "ticker"

	// MVP needed
	ORDER   = "order"
	ACCOUNT = "account"

	// MVP needed - Channels
	BTCUSDT = "btcusdt@depth5@100ms"
	ETHBTC  = "ethbtc@depth5@100ms"
	ETHUSDT = "ethusdt@depth5@100ms"
)
