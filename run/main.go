package main

import (
	"fmt"      // Package untuk formatting teks dan output
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/routes" // Package yang mungkin berisi definisi-definisi rute atau endpoint HTTP dalam aplikasi
)

func main() {
	// Menetapkan handler untuk route "/" ke fungsi URL dari package routes.
	http.HandleFunc("/", routes.URL)
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, nil)
}
