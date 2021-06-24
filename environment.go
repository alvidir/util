package util

import (
	"fmt"
	"os"
)

const (
	errEnvNotFound = "environment variable %s not found"
	errEnvIsEmpty  = "environment variable %s do not provide any content"
)

// LookupEnv looks up for the env variable key and returns it if exists, else err != nil
func LookupEnv(key string) (env string, err error) {
	var exists bool
	if env, exists = os.LookupEnv(key); !exists {
		err = fmt.Errorf(errEnvNotFound, key)
	}

	return
}

// LookupNempEnv looks up for the env variable key and returns it if exists and is not null, else err != nil
func LookupNempEnv(key string) (env string, err error) {
	var exists bool
	if env, exists = os.LookupEnv(key); !exists || len(env) == 0 {
		err = fmt.Errorf(errEnvNotFound, key)
	}

	return
}
