package config

import (
	"fmt"
	"os"
)

// CheckEnv looks up for the env variable key and returns it if exists, els err != nil
func CheckEnv(key string) (value string, err error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		err = fmt.Errorf(errEnvNotFound, key)
	}

	return
}
