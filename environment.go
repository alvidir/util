package util

import (
	"errors"
	"os"
)

const (
	ErrEnvNotFound = "environment variable not found"
	ErrEnvIsEmpty  = "environment variable do not provide any content"
)

// LookupEnv looks up for the env variable key and returns it if exists, else err != nil
func LookupEnv(key string) (env string, err error) {
	var exists bool
	if env, exists = os.LookupEnv(key); !exists {
		err = errors.New(ErrEnvNotFound)
	}

	return
}

// LookupNempEnv looks up for the env variable key and returns it if exists and is not null, else err != nil
func LookupNempEnv(key string) (env string, err error) {
	var exists bool
	if env, exists = os.LookupEnv(key); !exists || len(env) == 0 {
		err = errors.New(ErrEnvNotFound)
	}

	return
}
