// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/owulveryck/onnx-go"
// 	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
// 	"github.com/sugarme/tokenizer"
// 	"github.com/sugarme/tokenizer/decoder"
// 	"github.com/sugarme/tokenizer/model/bpe"
// 	"github.com/sugarme/tokenizer/pretokenizer"
// 	"github.com/sugarme/tokenizer/processor"
// 	"gorgonia.org/tensor"
// )

// func main() {
// 	http.HandleFunc("/", ChatPredictions)
// 	port := ":8080"
// 	fmt.Println("Server started at: http://localhost" + port)
// 	http.ListenAndServe(port, nil)
// }

// func ChatPredictions(w http.ResponseWriter, r *http.Request) {
// 	// tokenizer.CachedPath("./model", "tokenizer_config.json")
// 	// tokenizer.CachedDir = "./model"
// 	err := tokenizer.CleanCache()
// 	if err != nil {
// 		http.Error(w, "Failed to clean cache", http.StatusInternalServerError)
// 		return
// 	}
// 	// if err != nil {
// 	// 	http.Error(w, "Failed to load tokenizer", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	model2, err := bpe.NewBpeFromFiles("model/gpt2-vocab.json", "model/gpt2-merges.txt")
// 	if err != nil {
// 		http.Error(w, "Failed to load tokenizer", http.StatusInternalServerError)
// 		return
// 	}

// 	addPrefixSpace := true
// 	trimOffsets := true
// 	tk := tokenizer.NewTokenizer(model2)

// 	pretok := pretokenizer.NewByteLevel()
// 	pretok.SetAddPrefixSpace(addPrefixSpace)
// 	pretok.SetTrimOffsets(trimOffsets)
// 	tk.WithPreTokenizer(pretok)

// 	pprocessor := processor.NewByteLevelProcessing(pretok)
// 	tk.WithPostProcessor(pprocessor)

// 	bpeDecoder := decoder.NewBpeDecoder("Ġ")
// 	tk.WithDecoder(bpeDecoder)

// 	// bl := pretokenizer.

// 	// tk := tokenizer.NewTokenizer(model2)
// 	// Initialize the tokenizer
// 	// tokenizer := tokenizer.NewTokenizer(configFile)

// 	// bpe.NewBpeFromFiles("model/gpt2-vocab.json", "model/gpt2-merges.txt")
// 	// Chaced Directory

// 	// prefix := true
// 	// trim := true
// 	// tk := pretrained.GPT2(prefix, trim)
//     // Load ONNX model
//     modelData, err := os.ReadFile("./gpt2.onnx")
// 	if err != nil {
// 		http.Error(w, "Failed to load model file", http.StatusInternalServerError)
// 		return
// 	}

// 	// Initialize the Gorgonnx backend
// 	backend := gorgonnx.NewGraph()

// 	// Initialize the ONNX model with the Gorgonnx backend
// 	model := onnx.NewModel(backend)

// 	// Unmarshal the model
// 	err = model.UnmarshalBinary(modelData)
// 	if err != nil {
// 		http.Error(w, "Failed to unmarshal model", http.StatusInternalServerError)
// 		return
// 	}

// 	// Read input question from the request body
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
// 		return
// 	}
// 	var requestData map[string]string
// 	if err := json.Unmarshal(body, &requestData); err != nil {
// 		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
// 		return
// 	}
// 	question := requestData["question"]

// 	// Preprocess the question and create the input tensor for the model
// 	// Adjust preprocessing according to your model's requirements
// 	// inputShape := []int{1, len(question)} // Assuming your input shape; adjust as needed
// 	// inputTensor := tensor.New(tensor.Of(tensor.Float32), tensor.WithShape(inputShape...))

// 	// inputTensor to byte
// 	// inputBytes := []byte(question)
// 	inputSeq := tokenizer.NewInputSequence(question)

// 	encoded, err := tk.Encode(tokenizer.NewSingleEncodeInput(inputSeq), false)
// 	if err != nil {
// 		http.Error(w, "Failed to encode question", http.StatusInternalServerError)
// 		return
// 	}

// 	input := tensor.New(tensor.Of(tensor.Int64), tensor.WithShape(1, len(encoded.Ids)))

// 	for i, token := range encoded.Ids {
// 		input.SetAt(int64(token), i)
// 	}

// 	// onnxTensor, err := onnx.NewTensor(inputBytes)
// 	// if err != nil {
// 	// 	http.Error(w, "Failed to create ONNX tensor", http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// inputTensorData := inputTensor.Data().([]float32)

// 	// Fill inputTensorData with your input question data
// 	// for i, char := range question {
// 	// 	inputTensorData[i] = float32(char)
// 	// }

// 	var backend2 *gorgonnx.Graph

// 	// Set the input tensor in the model
// 	err = backend2.SetInput(0, input)
// 	if err != nil {
// 		http.Error(w, "Failed to set input tensor", http.StatusInternalServerError)
// 		return
// 	}

// 	// Run inference using the Gorgonnx backend
// 	// err = backend.Run()
// 	// if err != nil {
// 	// 	http.Error(w, "Failed to run inference " + err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// Get the output tensor
// 	outputTensor, err := model.GetOutputTensors()
// 	if err != nil {
// 		http.Error(w, "Failed to get output tensor", http.StatusInternalServerError)
// 		return
// 	}

// 	outputData := outputTensor[0].Data().([]float32)
// 	ints := make([]int, len(outputData))

//     for i, value := range outputData {
//         ints[i] = int(value)
//     }

// 	decodedOutput := tk.Decode(ints, true)

// 	output := fmt.Sprintf("Output: %v", decodedOutput)

// 	// Send the prediction result as the response
// 	w.Header().Set("Content-Type", "application/json")
// 	response := map[string]string{"prediction": output}
// 	json.NewEncoder(w).Encode(response)
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
	"gorgonia.org/tensor"
)

type ChatRequest struct {
	Input string `json:"input"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

var (
	model     *onnx.Model
	backend   *gorgonnx.Graph
	gpt2Token *tokenizer.Tokenizer
)

func initModel() {
	// Load the model data
	modelData, err := os.ReadFile("./gpt2.onnx")
	if err != nil {
		log.Fatalf("Failed to load model file: %v", err)
	}

	// Initialize the Gorgonnx backend
	backend = gorgonnx.NewGraph()

	// Initialize the ONNX model with the Gorgonnx backend
	model = onnx.NewModel(backend)

	// Unmarshal the model
	err = model.UnmarshalBinary(modelData)
	if err != nil {
		log.Fatalf("Failed to unmarshal model: %v", err)
	}
}

func initTokenizer() {
	gpt2Token = pretrained.GPT2(true, true)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var chatReq ChatRequest
	err := json.NewDecoder(r.Body).Decode(&chatReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	inputSeq := tokenizer.NewInputSequence(chatReq.Input)

	// Tokenize input
	encoded, err := gpt2Token.Encode(tokenizer.NewSingleEncodeInput(inputSeq), false)
	if err != nil {
		http.Error(w, "Failed to tokenize input", http.StatusInternalServerError)
		return
	}

	// Convert tokens to tensor
	inputTensor := tensor.New(tensor.Of(tensor.Int64), tensor.WithShape(1, len(encoded.Ids)))
	for i, token := range encoded.Ids {
		inputTensor.SetAt(int64(token), i)
	}

	// Set the input for the backend using onnx.Value
	input, err := onnx.NewTensor([]byte(inputTensor.DataOrder().String()))
	if err != nil {
		http.Error(w, "Failed to create input tensor", http.StatusInternalServerError)
		return
	}

	// Set the input tensor in the model
	err = model.SetInput(0, input)
	if err != nil {
		http.Error(w, "Failed to set input tensor", http.StatusInternalServerError)
		return
	}

	
	// Run the inference
	err = backend.Run()
	if err != nil {
		http.Error(w, "Failed to run inference", http.StatusInternalServerError)
		return
	}

	// Extract and process output
	output, err := model.GetOutputTensors()
	if err != nil {
		http.Error(w, "Failed to get output", http.StatusInternalServerError)
		return
	}

	// Decode output tokens
	decodedOutput := gpt2Token.Decode(output[0].Data().([]int), true)

	// Create response
	chatResp := ChatResponse{
		Response: decodedOutput,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatResp)
}

func main() {
	// Initialize the model and tokenizer
	initModel()
	initTokenizer()

	http.HandleFunc("/chat", chatHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
