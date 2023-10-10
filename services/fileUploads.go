// controllers/user_controller.go
package services

import (
	// ... (other imports)
	"io"
	"net/http"
	"os"
	"path/filepath"
)


// UploadFileHandler handles file uploads
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading form file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename (you can use a UUID or other strategies)
	fileName := handler.Filename

	// Save the file to the server
	targetPath := filepath.Join("uploads", fileName)
	f, err := os.Create(targetPath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}

	// Respond with a success message or redirect as needed
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}
