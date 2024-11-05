package go_binance_url_builder

import (
	"errors"
	"testing"
)

const (
	baseURLTestHTTP = "https://testnet.binance.vision/api"
	baseURLTestWSS  = "wss://ws-api.testnet.binance.vision/ws-api/v3"
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
