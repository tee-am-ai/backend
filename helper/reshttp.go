package helper

import (
	"encoding/json" // Package untuk encoding dan decoding data dalam format JSON
	"log"           // Package untuk logging, digunakan untuk mencatat informasi atau kesalahan ke konsol atau file log
	"net/http"      // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons
)

func ErrorResponse(respw http.ResponseWriter, req *http.Request, statusCode int, err, msg string) {
	resp := map[string]string{
		"error":   err,
		"message": msg,
	}
	WriteJSON(respw, statusCode, resp)
}

func WriteJSON(respw http.ResponseWriter, statusCode int, content any) {
	respw.Header().Set("Content-Type", "application/json")
	respw.WriteHeader(statusCode)
	respw.Write([]byte(Jsonstr(content)))
}

func Jsonstr(strc any) string {
	jsonData, err := json.Marshal(strc)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}
