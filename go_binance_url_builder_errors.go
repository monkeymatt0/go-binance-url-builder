package go_binance_url_builder

// This file will contain different errors implementing the error interface

import (
	"errors"
)

const (
	// Possible getHost function's errors
	MODE_ERROR        = "This mode does not exists"                     // Wrong mode is typed
	ORDER_URL_ERROR   = "Error while creating the order endpoint URL"   // Wrong mode is typed
	ACCOUNT_URL_ERROR = "Error while creating the account endpoint URL" // Wrong mode is typed
)

// MODE_ERROR can happen while calling getHost function
var ModeError = errors.New(MODE_ERROR)
var OrderURLError = errors.New(ORDER_URL_ERROR)
var AccountURLError = errors.New(ACCOUNT_URL_ERROR)
