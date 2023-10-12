package services

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) (string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return "", err
	}

	files := form.File["files"]
	if len(files) == 0 {
		return "", errors.New("No files uploaded")
	}

	var filePath string

	for key, file := range files {
		uploadedFilePath, err := saveFile(file, key)
		if err != nil {
			return "", err
		}

		filePath = uploadedFilePath
	}

	return filePath, nil
}

func saveFile(fileHeader *multipart.FileHeader, key int) (string, error) {
	ext := strings.Split(fileHeader.Filename, ".")
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	// Create a new file in the desired destination folder
	dstPath := filepath.Join("./destination", ext[0]+strconv.Itoa(key)+"."+ext[1])
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	return dstPath, nil
}
