package config

import (
	"os" // Mengimpor paket os untuk bekerja dengan lingkungan sistem operasi
)

// GetEnv adalah fungsi yang mengambil nilai dari variabel lingkungan berdasarkan nama yang diberikan.
func GetEnv(envName string) string {
	// Mengambil nilai dari variabel lingkungan dengan nama envName menggunakan os.Getenv.
	return os.Getenv(envName)
}
