package controller

import (
	"encoding/json" // Mengimpor paket encoding/json untuk bekerja dengan JSON
	"log"           // Mengimpor paket log untuk logging
	"net/http"      // Mengimpor paket net/http untuk menangani HTTP request dan response
	"net/url"       // Mengimpor paket net/url untuk bekerja dengan URL
	"strings"       // Mengimpor paket strings untuk manipulasi string
	"time"          // Mengimpor paket time untuk bekerja dengan waktu

	"github.com/go-resty/resty/v2"        // Mengimpor paket resty untuk membuat HTTP request dengan client yang lebih kaya fitur
	"github.com/tee-am-ai/backend/config" // Mengimpor package config dari aplikasi backend bagian config
	"github.com/tee-am-ai/backend/helper" // Mengimpor package helper dari aplikasi backend bagian helper
	"github.com/tee-am-ai/backend/model"  // Mengimpor package model dari aplikasi backend bagian model
)

// Chat adalah handler fungsi untuk menangani permintaan chat AI.
func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
	// Mendeklarasikan variabel chat sebagai tipe model.AIRequest
	var chat model.AIRequest

	// Mendekode body permintaan JSON ke dalam variabel chat
	err := json.NewDecoder(req.Body).Decode(&chat)
	if err != nil {
		// Jika terjadi kesalahan saat mendekode, kirimkan respons kesalahan 400 Bad Request
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	// Memastikan bahwa query tidak kosong
	if chat.Query == "" {
		// Jika query kosong, kirimkan respons kesalahan 400 Bad Request
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	// Inisialisasi klien untuk melakukan permintaan HTTP dengan library resty
	client := resty.New()

	// Mendapatkan URL dan token API dari environment dan parameter fungsi
	apiUrl := config.GetEnv("HUGGINGFACE_API_KEY")
	apiToken := "Bearer " + tokenmodel

	// Mendeklarasikan variabel untuk respons dan retry
	var response *resty.Response
	var retryCount int
	maxRetries := 5
	retryDelay := 20 * time.Second

	// Parsing URL API Hugging Face
	parsedURL, err := url.Parse(apiUrl)
	if err != nil {
		// Jika terjadi kesalahan saat parsing URL, kirimkan respons kesalahan 500 Internal Server Error
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model hugging face"+err.Error())
		return
	}

	// Mendapatkan nama model dari URL yang di-parse
	segments := strings.Split(parsedURL.Path, "/")
	modelName := strings.Join(segments[2:], "/")

	// Melakukan request ke Hugging Face API dengan mekanisme retry
	for retryCount < maxRetries {
		response, err = client.R().
			SetHeader("Authorization", apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(`{"inputs": "` + chat.Query + `"}`).
			Post(apiUrl)

		if err != nil {
			// Jika terjadi kesalahan saat membuat permintaan, hentikan program dengan pesan log
			log.Fatalf("Error making request: %v", err)
		}

		// Memeriksa status kode respons dari Hugging Face API
		if response.StatusCode() == http.StatusOK {
			break
		} else {
			// Menangani kasus ketika model masih dalam proses loading dengan retry
			var errorResponse map[string]interface{}
			err = json.Unmarshal(response.Body(), &errorResponse)
			if err == nil && errorResponse["error"] == "Model "+modelName+" is currently loading" {
				retryCount++
				time.Sleep(retryDelay)
				continue
			}
			// Jika terjadi kesalahan lain, kirimkan respons kesalahan 500 Internal Server Error
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
			return
		}
	}
	// Jika setelah maxRetries masih tidak berhasil, kirimkan respons kesalahan 500 Internal Server Error
	if response.StatusCode() != http.StatusOK {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
		return
	}

	// Dekode respons JSON dari Hugging Face API
	var data []map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		// Jika terjadi kesalahan saat mendekode respons, kirimkan respons kesalahan 500 Internal Server Error
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
	}

	// Memeriksa apakah ada data yang dihasilkan dari respons
	if len(data) > 0 {
		// Mengekstrak teks yang dihasilkan dari respons
		generatedText, ok := data[0]["generated_text"].(string)
		if !ok {
			// Jika terjadi kesalahan saat ekstraksi teks, kirimkan respons kesalahan 500 Internal Server Error
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
			return
		}
		// Mengirimkan respons dengan teks yang dihasilkan dalam format JSON
		helper.WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
	} else {
		// Jika tidak ada data yang dihasilkan, kirimkan respons kesalahan 500 Internal Server Error
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}
