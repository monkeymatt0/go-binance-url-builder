package go_binance_url_builder

import (
	"errors"
	"testing"
)

const (
	baseURLTestHTTP = "https://testnet.binance.vision"
	baseURLTestWSS  = "wss://testnet.binance.vision"
)

// Testing URL for HTTP request
func TestBaseURLSuccessHTTP(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, TEST)

	if err != nil {
		t.Errorf("Error when setting the host: %v\n", err)
	}

	if binanceURLBuilder.GetBaseURL() != baseURLTestHTTP {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestHTTP, binanceURLBuilder.GetBaseURL())
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
	err := binanceURLBuilder.New(true, TEST)

	if err != nil {
		t.Errorf("Error when setting the host: %v\n", err)
	}

	if binanceURLBuilder.GetBaseURL() != baseURLTestWSS {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestWSS, binanceURLBuilder.GetBaseURL())
	}
}

func TestBaseURLFailWSS(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	err := binanceURLBuilder.New(false, "")

	if !errors.Is(err, ModeError) {
		t.Errorf("This test should fail for ModeError")
	}
}
