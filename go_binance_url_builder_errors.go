package go_binance_url_builder

// This file will contain different errors implementing the error interface

import (
	"errors"
)

const (
	// Possible getHost function's errors
	MODE_ERROR = "This mode does not exists" // Wrong mode is typed
)

// MODE_ERROR can happen while calling getHost function
var ModeError = errors.New(MODE_ERROR)
