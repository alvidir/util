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

// LookupAllEnv looks up for all the env keys and returns a map of values for each key, if exists, else err != nil
func LookupAllEnv(keys ...string) (env map[string]string, err error) {
	env = make(map[string]string)
	for _, key := range keys {
		var exists bool
		if env[key], exists = os.LookupEnv(key); !exists {
			err = fmt.Errorf(errEnvNotFound, key)
			return
		}
	}

	return
}

// LookupNempEnv looks up for all the env keys and returns a map of values for each key, if exists and not empty, else err != nil
func LookupNempEnv(keys ...string) (env map[string]string, err error) {
	env = make(map[string]string, len(keys))
	for _, key := range keys {
		var exists bool
		if env[key], exists = os.LookupEnv(key); !exists {
			err = fmt.Errorf(errEnvNotFound, key)
			return
		}

		if len(env[key]) == 0 {
			delete(env, key)
			err = fmt.Errorf(errEnvIsEmpty, key)
			return
		}
	}

	return
}
