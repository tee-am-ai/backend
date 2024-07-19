package main

import (
	"fmt"      // Package untuk formatting teks dan output
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/routes" // Package yang mungkin berisi definisi-definisi rute atau endpoint HTTP dalam aplikasi
)

// Fungsi main() adalah titik masuk utama dari aplikasi Go.
// Fungsi ini akan dijalankan saat aplikasi dimulai.
func main() {
	// Menetapkan handler untuk rute "/" dengan fungsi routes.URL
	// Fungsi routes.URL akan menangani semua permintaan HTTP yang masuk ke root path
	http.HandleFunc("/", routes.URL)

	// Menentukan port yang akan digunakan untuk server HTTP
	port := ":8080"

	// Mencetak pesan ke konsol yang menunjukkan bahwa server telah dimulai
	// dan menunjukkan URL tempat server dapat diakses
	fmt.Println("Server started at: http://localhost" + port)

	// Memulai server HTTP pada port yang telah ditentukan
	// http.ListenAndServe akan mendengarkan permintaan HTTP dan meneruskan ke handler yang telah ditentukan
	// Dalam hal ini, handler adalah fungsi routes.URL
	http.ListenAndServe(port, nil)
}
