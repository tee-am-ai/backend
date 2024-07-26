package config

import (
	"os"
)

func GetEnv(envName string) string {
	return os.Getenv(envName)
}
