# go-binance-url-builder

This package is primarly used for another private projects, so it will implement first the URL needed in that project.

The goal of this package is to return a string with the correct endpoint to call (for now).

### Usage

To properly create a BinanceURLBuilder you have to create it like this:

```

     binanceURLBuilder := &BinanceURLBuilder{}
     err := binanceURLBuilder.New(false, TEST)

```

In this case you will create a BinanceURLBuilder made for HTTP URL since this is the New function:

```
func (bub BinanceURLBuilder) New(isWebSocketStream bool, mode string) error {
     bub.SetProtocol(isWebSocket)
     err := bub.setHost(mode)
     if err != nil {
          return err
     }
     bub.setVersion()
     return nil
}
```

It takes isWebSocketStream => if true then it will be create the base url for the web socket stream
It takes isWebSocketStream => if false then it will return HTTP URL

It takes mode, possible values for mode:

```
     PRODUCTION = "production"
     PRODUCTION_WS = "production_ws"
     PRODUCTION_WSS = "production_wss"
     TEST = "test"
     TEST_WS = "test_ws"
     TEST_WSS = "test_wss"
```

For now \*\_WS are not handled

### Function that will be used (for now)

```
     func (bub *BinanceURLBuilder) GetBaseURLWSS() string
```

This function will return the following string (URL):

```
     wss://stream.binance.com:9443/ws (if in PRODUCTION_WSS)
     wss://stream.testnet.binance.vision:9443/ws (if in TEST_WSS)
```

---

```
     func (bub *BinanceURLBuilder) GetAccountURL() string
```

This function will return

```
     https://api.binance.com/api/v3/account (if in PRODUCTION)
     https://testnet.binance.vision/api/v3/account (if in TEST)
```

---

```
     func (bub *BinanceURLBuilder) GetOrderURL() string
```

This function will return

```
     https://api.binance.com/api/v3/order (if in PRODUCTION)
     https://testnet.binance.vision/api/v3/order (if in TEST)
```

For now these are the MVP url needed
