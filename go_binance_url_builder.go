package go_binance_url_builder

import "strings"

// @todo : This will be refactored in future

// For now the only goal of this structure is to be used
// as a single type of builder (HTTP or WEB SOCKET STREAM)
type BinanceURLBuilder struct {
	Schema       string
	Host         string
	version      string
	pathFragment string
}

func (bub *BinanceURLBuilder) New(isWebSocketStream bool, mode string) error {
	bub.SetProtocol(isWebSocketStream)
	err := bub.setHost(mode)
	if err != nil {
		return err
	}
	bub.setVersion()
	return nil
}

func (bub *BinanceURLBuilder) SetProtocol(isWebSocketStream bool) {
	if isWebSocketStream {
		bub.Schema = strings.Join([]string{SECURE_WEB_SOCKET, "://"}, "")
	} else {
		bub.Schema = strings.Join([]string{SECURE_HTTP, "://"}, "")
	}
}

// This will return me if the I should connect to the production os test host
func (bub *BinanceURLBuilder) setHost(mode string) error {
	switch mode {
	case PRODUCTION:
		bub.Host = strings.Join([]string{PRODUCTION_HOST}, "/")
		return nil
	case TEST:
		bub.Host = strings.Join([]string{TEST_HOST}, "/")
		return nil
	case TEST_WSS:
		bub.Host = strings.Join([]string{TEST_WSS_HOST}, "/")
		return nil
	default:
		return ModeError
	}
}

func (bub *BinanceURLBuilder) GetBaseURLHTTP() string {
	return strings.Join([]string{strings.Join([]string{bub.Schema, bub.Host}, ""), API}, "/")
}

func (bub *BinanceURLBuilder) GetBaseURLWSS() string {
	return strings.Join([]string{strings.Join([]string{bub.Schema, bub.Host}, ""), WSS_API}, "/")
}

// This will always set to v3 since is the most realiable and updated for now
func (bub *BinanceURLBuilder) setVersion() {
	bub.version = V3
}

func (bub *BinanceURLBuilder) GetAccountURL() string {
	baseURL := bub.GetBaseURLHTTP()
	accountURLSlice := []string{baseURL, bub.version, ACCOUNT}
	return strings.Join(accountURLSlice, "/")
}

func (bub *BinanceURLBuilder) GetOrderURL() string {
	baseURL := bub.GetBaseURLHTTP()
	accountURLSlice := []string{baseURL, bub.version, ORDER}
	return strings.Join(accountURLSlice, "/")
}
