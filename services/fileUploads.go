// services/upload.go
package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleUpload(c *gin.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File["files"]
	var filePaths []string

	for _, file := range files {
		filePath, err := saveFile(file)
		if err != nil {
			return nil, err
		}
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

func saveFile(fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)

	// Specify the destination directory as an absolute path based on the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fmt.Printf("File currentDir to: %s\n", currentDir)

	// Define the relative destination directory (relative to the current working directory)
	destinationDir := "destination"

	// Ensure the destination directory exists
	fullDestinationDir := filepath.Join(currentDir, destinationDir)
	if _, err := os.Stat(fullDestinationDir); os.IsNotExist(err) {
		if err := os.MkdirAll(fullDestinationDir, os.ModePerm); err != nil {
			return "", err
		}
	}

	dstPath := filepath.Join(destinationDir, uuid.New().String()+ext)

	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fullDstPath := filepath.Join(currentDir, dstPath)

	dst, err := os.Create(fullDstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return dstPath, nil
}
