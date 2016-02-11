package variables

import (
	"os"
)

var (
	// EnvEnableAuth gets EnableAuth from an env variable
	EnvEnableAuth = "ENABLE_AUTH"
	// EnableAuth lets whoever uses this service turn off auth for testing
	EnableAuth bool
	// EnableAuthDefault  lets whoever uses this service turn off auth for testing programaticly
	EnableAuthDefault = true
)

func init() {
	if os.Getenv(EnvEnableAuth) == "" {
		EnableAuth = EnableAuthDefault
	} else if os.Getenv(EnvEnableAuth) == "true" {
		EnableAuth = true
	} else if os.Getenv(EnvEnableAuth) == "false" {
		EnableAuth = false
	}
}
