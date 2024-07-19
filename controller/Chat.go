// package controller

// import (
// 	"encoding/json" // Mengimpor paket encoding/json untuk bekerja dengan JSON
// 	"log"           // Mengimpor paket log untuk logging
// 	"net/http"      // Mengimpor paket net/http untuk menangani HTTP request dan response
// 	"net/url"       // Mengimpor paket net/url untuk bekerja dengan URL
// 	"strings"       // Mengimpor paket strings untuk manipulasi string
// 	"time"          // Mengimpor paket time untuk bekerja dengan waktu

// 	"github.com/go-resty/resty/v2"        // Mengimpor paket resty untuk membuat HTTP request dengan client yang lebih kaya fitur
// 	"github.com/tee-am-ai/backend/config" // Mengimpor package config dari aplikasi backend bagian config
// 	"github.com/tee-am-ai/backend/helper" // Mengimpor package helper dari aplikasi backend bagian helper
// 	"github.com/tee-am-ai/backend/model"  // Mengimpor package model dari aplikasi backend bagian model
// )

// func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
// 	var chat model.AIRequest
// 	err := json.NewDecoder(req.Body).Decode(&chat)
// 	if err != nil {
// 		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
// 		return
// 	}

// 	// Memastikan bahwa query tidak kosong
// 	if chat.Query == "" {
// 		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
// 		return
// 	}

// 	// Inisialisasi klien untuk melakukan permintaan HTTP dengan library resty
// 	client := resty.New()

// 	// Mendapatkan URL dan token API dari environment dan parameter fungsi
// 	apiUrl := config.GetEnv("HUGGINGFACE_API_KEY")
// 	apiToken := "Bearer " + tokenmodel

// 	var response *resty.Response
// 	var retryCount int
// 	maxRetries := 5
// 	retryDelay := 20 * time.Second

// 	// Parsing URL API Hugging Face
// 	parsedURL, err := url.Parse(apiUrl)
// 	if err != nil {
// 		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model hugging face"+err.Error())
// 		return
// 	}

// 	segments := strings.Split(parsedURL.Path, "/")
// 	modelName := strings.Join(segments[2:], "/")

// 	// Melakukan request ke Hugging Face API dengan mekanisme retry
// 	for retryCount < maxRetries {
// 		response, err = client.R().
// 			SetHeader("Authorization", apiToken).
// 			SetHeader("Content-Type", "application/json").
// 			SetBody(`{"inputs": "` + chat.Query + `"}`).
// 			Post(apiUrl)

// 		if err != nil {
// 			log.Fatalf("Error making request: %v", err)
// 		}

// 		// Memeriksa status kode respons dari Hugging Face API
// 		if response.StatusCode() == http.StatusOK {
// 			break
// 		} else {
// 			var errorResponse map[string]interface{}
// 			err = json.Unmarshal(response.Body(), &errorResponse)
// 			if err == nil && errorResponse["error"] == "Model "+modelName+" is currently loading" {
// 				retryCount++
// 				time.Sleep(retryDelay)
// 				continue
// 			}
// 			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
// 			return
// 		}
// 	}
// 	if response.StatusCode() != http.StatusOK {
// 		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
// 		return
// 	}

// 	// Dekode respons JSON dari Hugging Face API
// 	var data []map[string]interface{}
// 	err = json.Unmarshal(response.Body(), &data)
// 	if err != nil {
// 		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
// 		return
// 	}

// 	// Memeriksa apakah ada data yang dihasilkan dari respons
// 	if len(data) > 0 {
// 		generatedText, ok := data[0]["generated_text"].(string)
// 		if !ok {
// 			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
// 			return
// 		}
// 		helper.WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
// 	} else {
// 		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
// 	}
// }
