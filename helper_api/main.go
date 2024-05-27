package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/debian", func(w http.ResponseWriter, r *http.Request) {
		// Read the Bash script file
		script, err := os.ReadFile("debian.sh")
		if err != nil {
			http.Error(w, "Failed to read Bash script", http.StatusInternalServerError)
			return
		}

		// Set content type header
		w.Header().Set("Content-Type", "text/plain")

		// Serve the Bash script
		_, err = w.Write(script)
		if err != nil {
			http.Error(w, "Failed to serve Bash script", http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
