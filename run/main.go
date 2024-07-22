package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/sugarme/tokenizer/pretrained"
	"gorgonia.org/tensor"
)

func main() {
	http.HandleFunc("/", ChatPredictions)
	port := ":8080"
	fmt.Println("Server started at: http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func ChatPredictions(w http.ResponseWriter, r *http.Request) {
	// configFile, err := tokenizer.CachedPath("./", "tokenizer_config.json")
	// if err != nil {
	// 	http.Error(w, "Failed to load tokenizer", http.StatusInternalServerError)
	// 	return
	// }
	tokenizer, err := pretrained.FromFile("./tokenizer_config.json")
	if err != nil {
		http.Error(w, "Failed to load tokenizer", http.StatusInternalServerError)
		return
	}
    // Load ONNX model
    modelData, err := os.ReadFile("./gpt2.onnx")
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
	// inputShape := []int{1, len(question)} // Assuming your input shape; adjust as needed
	// inputTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(inputShape...))

	// inputTensor to byte
	// inputBytes := []byte(question)

	encoded, err := tokenizer.EncodeSingle(question)
	if err != nil {
		http.Error(w, "Failed to encode question", http.StatusInternalServerError)
		return
	}

	input := tensor.New(tensor.Of(tensor.Int64), tensor.WithShape(1, len(encoded.Ids)))

	for i, token := range encoded.Ids {
		input.SetAt(int64(token), i)
	}	
	

	// onnxTensor, err := onnx.NewTensor(inputBytes)
	// if err != nil {
	// 	http.Error(w, "Failed to create ONNX tensor", http.StatusInternalServerError)
	// 	return
	// }
	// inputTensorData := inputTensor.Data().([]float32)

	// Fill inputTensorData with your input question data
	// for i, char := range question {
	// 	inputTensorData[i] = float32(char)
	// }

	// Set the input tensor in the model
	err = model.SetInput(0, input)
	if err != nil {
		http.Error(w, "Failed to set input tensor", http.StatusInternalServerError)
		return
	}

	// Run inference using the Gorgonnx backend
	err = backend.Run()
	if err != nil {
		http.Error(w, "Failed to run inference " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the output tensor
	outputTensor, err := model.GetOutputTensors()
	if err != nil {
		http.Error(w, "Failed to get output tensor", http.StatusInternalServerError)
		return
	}

	output := fmt.Sprintf("Output: %v", outputTensor)

	// Send the prediction result as the response
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"prediction": output}
	json.NewEncoder(w).Encode(response)
}
