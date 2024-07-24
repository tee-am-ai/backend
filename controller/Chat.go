package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tee-am-ai/backend/config"
	"github.com/tee-am-ai/backend/helper"
	"github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Chat(db *mongo.Database, respw http.ResponseWriter, req *http.Request, tokenmodel, PASETOPUBLICKEYENV string) {
	tokenstring := req.Header.Get("Login")
	payload, err := helper.Decode(PASETOPUBLICKEYENV, tokenstring)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error decoding token "+err.Error())
		return
	}

	// Parse request body
	var chat model.AIRequest

	err = json.NewDecoder(req.Body).Decode(&chat)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if chat.Query == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	client := resty.New()

	// Hugging Face API URL dan token
	apiUrl := config.GetEnv("HUGGINGFACE_API_KEY")
	apiToken := "Bearer " + tokenmodel

	var response *resty.Response
	var retryCount int
	maxRetries := 5
	retryDelay := 20 * time.Second
	
	parsedURL, err := url.Parse(apiUrl)

	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model hugging face"+err.Error())
        return
    }

	segments := strings.Split(parsedURL.Path, "/")

	modelName := strings.Join(segments[2:], "/")

	// Request ke Hugging Face API
	for retryCount < maxRetries {
		response, err = client.R().
			SetHeader("Authorization", apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(`{"inputs": "` + chat.Query + `"}`).
			Post(apiUrl)

		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}

		if response.StatusCode() == http.StatusOK {
			break
		} else {
			var errorResponse map[string]interface{}
			err = json.Unmarshal(response.Body(), &errorResponse)
			if err == nil && errorResponse["error"] == "Model " + modelName + " is currently loading" {
				retryCount++
				time.Sleep(retryDelay)
				continue
			}
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
			return
		}
	}

	if response.StatusCode() != 200 {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Errorr", "error from Hugging Face API "+string(response.Body()))
		return
	}

	var data []map[string]interface{}

	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
	}

	if len(data) > 0 {
		generatedText, ok := data[0]["generated_text"].(string)
		if !ok {
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
			return
		}
		helper.WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
	} else {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}

func Chat2(respw http.ResponseWriter, req *http.Request) {
	url := "http://localhost:5000/generate"

	var chat model.AIRequest

	err := json.NewDecoder(req.Body).Decode(&chat)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if chat.Query == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}


    // Buat payload JSON
    payload := map[string]string{"question": chat.Query}
    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing request body "+err.Error())
		return
    }

    // Buat request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
    if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error making request "+err.Error())
		return
    }
    defer resp.Body.Close()

    // Baca response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error reading response body "+err.Error())
		return
    }


    // Parse response
    var result map[string]string
    if err := json.Unmarshal(body, &result); err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
    }

	returndata := map[string]string{"answer": result["answer"]}
	helper.WriteJSON(respw, http.StatusOK, returndata)
}