package main

import (
	"fmt"      // Package untuk formatting teks dan output
	"net/http" // Package untuk melakukan operasi-operasi terkait HTTP seperti membuat server, mengirim permintaan, dan menerima respons HTTP

	"github.com/tee-am-ai/backend/routes" // Package yang mungkin berisi definisi-definisi rute atau endpoint HTTP dalam aplikasi
)

func main() {
	// Register the "/" route to the URL handler defined in routes package.
	http.HandleFunc("/", routes.URL)

	// Define the port for the server to listen on.
	port := ":8080"

	// Print a message indicating the server has started.
	fmt.Println("Server started at: http://localhost" + port)

	// Start the HTTP server and listen on the specified port.
	http.ListenAndServe(port, nil)
}
