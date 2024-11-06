package go_binance_url_builder

import (
	"errors"
	"testing"
)

const (
	baseURLTestHTTP        = "https://testnet.binance.vision/api"
	baseURLTestHTTPOrder   = "https://testnet.binance.vision/api/v3/order"   // Place order
	baseURLTestHTTPAccount = "https://testnet.binance.vision/api/v3/account" // get account info like balances
	baseURLTestWSS         = "wss://stream.testnet.binance.vision:9443/ws"

	baseURLProductionHTTP        = "https://api.binance.com/api"
	baseURLProductionHTTPOrder   = "https://api.binance.com/api/v3/order"
	baseURLProductionHTTPAccount = "https://api.binance.com/api/v3/account"
	baseURLProductionWSS         = "wss://stream.binance.com:9443/ws"
)

// Testnet test
// Testing URL for HTTP request
func TestBaseURLSuccessHTTP(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, TEST)

	if err != nil {
		t.Errorf("Error when setting the host: %v\n", err)
	}

	if binanceURLBuilder.GetBaseURLHTTP() != baseURLTestHTTP {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestHTTP, binanceURLBuilder.GetBaseURLHTTP())
	}
}

func TestBaseURLFailHTTP(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, "")

	if !errors.Is(err, ModeError) {
		t.Errorf("This test should fail for ModeError: %s\n", err)
	}
}

// Testing base URL for web socket
func TestBaseURLSuccessWSS(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(true, TEST_WSS)

	if err != nil {
		t.Errorf("Error when setting the host: %v\n", err)
	}

	if binanceURLBuilder.GetBaseURLWSS() != baseURLTestWSS {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestWSS, binanceURLBuilder.GetBaseURLWSS())
	}

}

func TestBaseURLFailWSS(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, "")

	if !errors.Is(err, ModeError) {
		t.Errorf("This test should fail for ModeError")
	}
}

func TestOrderEndpointURL(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, TEST)
	if err != nil {
		t.Errorf("Error when setting the host(TEST): %s\n", err)
	}
	if binanceURLBuilder.GetOrderURL() != baseURLTestHTTPOrder {
		t.Errorf("Error(TEST): %s\n Got: %s\n Want: %s\n", OrderURLError, binanceURLBuilder.GetOrderURL(), baseURLTestHTTPOrder)
	}

	err2 := binanceURLBuilder.New(false, PRODUCTION)
	if binanceURLBuilder.GetOrderURL() != baseURLProductionHTTPOrder {
		t.Errorf("Error(PRODUCTION): %s\n Got: %s\n Want: %s\n", OrderURLError, binanceURLBuilder.GetOrderURL(), baseURLProductionHTTP)
	}

	if err2 != nil {
		t.Errorf("Error when setting the host(PRODUCTION): %s\n", err2)
	}

}

func TestAccountEndpointURL(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, TEST)
	if err != nil {
		t.Errorf("Error when setting the host(TEST): %s\n", err)
	}
	if binanceURLBuilder.GetAccountURL() != baseURLTestHTTPAccount {
		t.Errorf("Error(TEST): %s\n", AccountURLError)
	}

	err2 := binanceURLBuilder.New(false, PRODUCTION)
	if binanceURLBuilder.GetAccountURL() != baseURLProductionHTTPAccount {
		t.Errorf("Error(PRODUCTION): %s\n", AccountURLError)
	}

	if err2 != nil {
		t.Errorf("Error when setting the host(PRODUCTION): %s\n", err2)
	}

}
