package go_binance_url_builder

import "testing"

const (
	baseURLTestHTTP = "https://testnet.binance.vision"
	baseURLTestWSS  = "wss://testnet.binance.vision"
)

func TestBaseURLSuccessHTTP(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	binanceURLBuilder.New(false, TEST)

	if binanceURLBuilder.GetBaseURL() != baseURLTestHTTP {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestHTTP, binanceURLBuilder.GetBaseURL())
	}
}

func TestBaseURLFailHTTP(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	binanceURLBuilder.New(false, "")

	if binanceURLBuilder.GetBaseURL() == baseURLTestHTTP {
		t.Errorf("BaseURL should not build want: %s, got: %s", baseURLTestHTTP, binanceURLBuilder.GetBaseURL())
	}
}
func TestBaseURLSuccessWSS(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	binanceURLBuilder.New(true, TEST)

	if binanceURLBuilder.GetBaseURL() != baseURLTestWSS {
		t.Errorf("BaseURL has not been well built want: %s, got: %s", baseURLTestWSS, binanceURLBuilder.GetBaseURL())
	}
}

func TestBaseURLFailWSS(t *testing.T) {
	binanceURLBuilder := &BinanceURLBuilder{}
	binanceURLBuilder.New(true, "")

	if binanceURLBuilder.GetBaseURL() == baseURLTestWSS {
		t.Errorf("BaseURL should not build want: %s, got: %s", baseURLTestWSS, binanceURLBuilder.GetBaseURL())
	}
}
