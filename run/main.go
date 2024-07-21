package main

import (
	"fmt"
	"net/http"
	"os"

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
}
