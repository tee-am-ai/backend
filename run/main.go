package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/tee-am-ai/backend/routes"
)

func main() {
	http.HandleFunc("/", routes.URL)
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func ChatPredictions(w http.ResponseWriter, r *http.Request) {
    // Load ONNX model
    modelData, err := os.ReadFile("path/to/your/model.onnx")
	if err != nil {
		http.Error(w, "Failed to load model file", http.StatusInternalServerError)
		return
	}

	// Initialize the Gorgonnx backend
	backend := gorgonnx.NewGraph()

	// Initialize the ONNX model with the Gorgonnx backend
	model := onnx.NewModel(backend)

	// Unmarshal the model
	err = model.UnmarshalBinary(modelData)
	if err != nil {
		http.Error(w, "Failed to unmarshal model", http.StatusInternalServerError)
		return
	}

	// Read input question from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	question := requestData["question"]

	// Preprocess the question and create the input tensor for the model
	// Adjust preprocessing according to your model's requirements
	inputShape := []int{1, len(question)} // Assuming your input shape; adjust as needed
}
