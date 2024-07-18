package config

import (
	"os" // Mengimpor paket os untuk bekerja dengan lingkungan sistem operasi
	// "github.com/joho/godotenv" // Mengimpor paket godotenv (dikomentari)
)

// Fungsi GetEnv mengambil nilai dari environment variable berdasarkan nama yang diberikan
func GetEnv(envName string) string {
	// envFile, _ := godotenv.Read("../.env")
	// return envFile[envName] 
	return os.Getenv(envName) // Mengembalikan nilai dari environment variable menggunakan os.Getenv
}
