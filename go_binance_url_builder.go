package go_binance_url_builder

import "strings"

type BinanceURLBuilder struct {
	protocol string
	host     string
}

func (bub *BinanceURLBuilder) New(isWebSocket bool, mode string) {
	bub.SetProtocol(isWebSocket)
	bub.setHost(mode)
}

func (bub *BinanceURLBuilder) SetProtocol(isWebSocket bool) {
	if isWebSocket {
		bub.protocol = strings.Join([]string{SECURE_WEB_SOCKET, "://"}, "")
	} else {
		bub.protocol = strings.Join([]string{SECURE_HTTP, "://"}, "")
	}
}

// This will return me if the I should connect to the production os test host
func (bub *BinanceURLBuilder) setHost(mode string) error {
	switch mode {
	case PRODUCTION:
		bub.host = strings.Join([]string{PRODUCTION_HOST}, "/")
		return nil
	case TEST:
		bub.host = strings.Join([]string{TEST_HOST}, "/")
		return nil
	case "default":
		return &ModeError{}
	}

	return &GetHostSwitchFailedError{} // This should never happen
}

func (bub *BinanceURLBuilder) GetBaseURL() string {
	return strings.Join([]string{bub.protocol, bub.host}, "")
}