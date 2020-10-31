package config

import (
	"fmt"
	"os"
)

// CheckEnv looks up for the env variable key and returns it if exists, else err != nil
func CheckEnv(key string) (value string, err error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		err = fmt.Errorf(errEnvNotFound, key)
	}

	return
}

// CheckAllEnv looks up for all the env keys and returns an slice of its values, if exists, else err != nil
func CheckAllEnv(keys ...string) (values []string, err error) {
	values = make([]string, len(keys))
	for index, key := range keys {
		var exists bool
		if values[index], exists = os.LookupEnv(key); !exists {
			err = fmt.Errorf(errEnvNotFound, key)
			return
		}
	}

	return
}
