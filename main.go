package teeamai

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions" // Package yang mendukung pengembangan fungsi-fungsi cloud di Google Cloud Platform
	"github.com/tee-am-ai/backend/routes"                             // Package yang mungkin berisi definisi-definisi rute atau endpoint HTTP dalam aplikasi
)

func init() {
	// Menggunakan fungsi HTTP dari functions-framework-go untuk mendaftarkan handler HTTP.
	// "WebHook" adalah nama fungsi yang digunakan untuk mendaftarkan handler.
	// routes.URL adalah handler HTTP yang akan dipanggil ketika ada permintaan ke "WebHook".
	functions.HTTP("WebHook", routes.URL)
}
