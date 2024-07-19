package config

import (
	"os" // Mengimpor paket os untuk bekerja dengan lingkungan sistem operasi
)

func GetEnv(envName string) string {
	return os.Getenv(envName)
}
