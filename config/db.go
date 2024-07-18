package config

import (
	"github.com/tee-am-ai/backend/helper" // Package yang mungkin berisi fungsi-fungsi bantuan terkait konfigurasi atau pengaturan
)

// MongoString adalah variabel untuk menyimpan string koneksi MongoDB yang didapatkan dari environment variable "MONGOSTRING".
// Nilai ini diambil dari fungsi GetEnv yang diasumsikan terdapat di dalam package helper.
var MongoString string = GetEnv("MONGOSTRING")

// mongoinfo adalah variabel untuk menyimpan informasi koneksi MongoDB seperti string koneksi dan nama database.
// Ini menggunakan struktur DBInfo yang mungkin didefinisikan di dalam package helper.
var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_teeamai",
}

// Mongoconn adalah variabel untuk menyimpan hasil koneksi MongoDB.
// ErrorMongoconn akan menyimpan nilai error jika terjadi kesalahan saat melakukan koneksi.
// Koneksi MongoDB diinisialisasi menggunakan fungsi MongoConnect yang mungkin didefinisikan di dalam package helper.
var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
