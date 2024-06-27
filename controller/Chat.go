package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/tee-am-ai/backend/helper"
	"github.com/tee-am-ai/backend/model"
)

func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
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

	client := resty.New()

    // Hugging Face API URL dan token
    apiUrl := "https://api-inference.huggingface.co/models/fatwaff/gpt2-model"
    apiToken := "Bearer " + tokenmodel

	response, err := client.R().
		SetHeader("Authorization", apiToken).
		SetHeader("Content-Type", "application/json").
		SetBody(`{"inputs": "`+chat.Query+`"}`).
		Post(apiUrl)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var data []map[string]string

	err = json.Unmarshal([]byte(response.String()), &data)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
	}

	if len(data) > 0 {
		helper.WriteJSON(respw, http.StatusOK, data[0])
	} else {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : response")
	}
}