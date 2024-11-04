package go_binance_url_builder

// This file will contain different errors implementing the error interface

import "fmt"

const (
	// Possible getHost function's errors
	MODE_ERROR             = "This mode does not exists"            // Wrong mode is typed
	GET_HOST_SWITCH_FAILED = "switch in getHost func failed... :\\" // In some way switch failed, should never happen
)

// MODE_ERROR can happen while calling getHost function
type ModeError struct{}

// Errors function implementations
func (ce *ModeError) Error() string {
	return fmt.Sprintf("----------\n%s\n\n----------\n", MODE_ERROR)
}

// GET_HOST_SWITCH_FAILED can happen if the switch in the getHost function in some ways fails
type GetHostSwitchFailedError struct{}

func (ghsfe *GetHostSwitchFailedError) Error() string {
	return fmt.Sprintf("----------\n%s\n\n----------\n", GET_HOST_SWITCH_FAILED)
}
