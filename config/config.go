package config

import (
	"os" // Mengimpor paket os untuk bekerja dengan lingkungan sistem operasi
)

// Fungsi GetEnv mengambil nilai dari environment variable berdasarkan nama yang diberikan
func GetEnv(envName string) string {
	return os.Getenv(envName) // Mengembalikan nilai dari environment variable menggunakan os.Getenv
}
