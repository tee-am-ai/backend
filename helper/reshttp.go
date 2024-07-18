package helper

import (
	"encoding/json" // Package untuk encoding dan decoding data dalam format JSON
	"log"           // Package untuk logging, digunakan untuk mencatat informasi atau kesalahan ke konsol atau file log
	"net/http"      // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons
)

// ErrorResponse adalah fungsi untuk mengirim respons JSON yang berisi pesan kesalahan ke klien.
// Fungsi ini memerlukan parameter:
// - respw http.ResponseWriter: objek untuk menulis respons HTTP.
// - req *http.Request: permintaan HTTP yang diterima.
// - statusCode int: kode status HTTP yang akan dikirimkan (misalnya 400 untuk Bad Request).
// - err string: jenis kesalahan yang terjadi.
// - msg string: pesan terkait dengan kesalahan tersebut.
func ErrorResponse(respw http.ResponseWriter, req *http.Request, statusCode int, err, msg string) {
	// Menyiapkan respons JSON yang berisi informasi kesalahan
	resp := map[string]string{
		"error":   err,
		"message": msg,
	}

	// Memanggil fungsi WriteJSON untuk menulis respons dalam format JSON dengan kode status yang sesuai
	WriteJSON(respw, statusCode, resp)
}

func WriteJSON(respw http.ResponseWriter, statusCode int, content any) {
	// Mengatur header untuk menetapkan jenis konten sebagai application/json.
	respw.Header().Set("Content-Type", "application/json")

	// Menetapkan status kode HTTP untuk respons.
	respw.WriteHeader(statusCode)

	// Mengubah konten ke dalam bentuk string JSON menggunakan fungsi Jsonstr.
	jsonContent := Jsonstr(content)

	// Menulis respons dalam bentuk byte array ke http.ResponseWriter.
	respw.Write([]byte(jsonContent))
}


func Jsonstr(strc any) string {
	jsonData, err := json.Marshal(strc)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}
