package utils

import (
	"errors"
	"fmt"
	"os"
)

func ValidString(key string, config string) (string, error) {
	if tmp := os.Getenv(key); tmp != "" {
		return tmp, nil
	}
	if config != "" {
		return config, nil
	}
	return "", errors.New(fmt.Sprintf("`%s` no configured", key))
}
