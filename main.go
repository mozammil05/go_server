// package main

// import (
// 	"fmt"
// 	"log"
// 	"my-auth-app/routes"
// 	"my-auth-app/utils"
// 	"os"

// 	// _ "my-auth-app/docs"

// 	"github.com/joho/godotenv"
// )

// // @title My Auth App API
// // @version 1.0
// // @description This is the API for My Auth App.
// // @host localhost:8080
// // @BasePath /api/v1
// // @schemes http
// // @schemes https
// func main() {
// 	// Load environment variables from the .env file
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Get the JWT secret key from the environment variables
// 	jwtSecret := os.Getenv("JWT")

// 	// Initialize MongoDB connection
// 	db := utils.InitDB()
// 	defer db.Disconnect()

// 	// Check if the MongoDB connection was successful
// 	if db.Client != nil {
// 		fmt.Println("MongoDB connected successfully")
// 	}

// 	// Get the port number from the environment variables
// 	port := os.Getenv("PORT")

// 	// Use a default port if the environment variable is not set
// 	if port == "" {
// 		port = "8080" // Default port
// 	}

// 	// Create a new router
// 	router := routes.NewRouter(db, jwtSecret)

// 	// Serve Swagger UI
// 	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

//		// Start your server on the specified port
//		router.Run(":" + port)
//	}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.POST("/upload", handleUpload)

// 	if err := r.Run(":8000"); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func handleUpload(c *gin.Context) {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Convert the form data to JSON and print it
// 	formData := make(map[string]interface{})
// 	// formData["values"] = form.Value
// 	formData["files"] = form.File
// 	fmt.Printf("Form Data: %+v\n", formData)

// 	files := form.File["files"]
// 	for key, file := range files {
// 		err := saveFile(file, key)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
// }

// func saveFile(fileHeader *multipart.FileHeader, key int) error {
// 	fmt.Println(fileHeader.Filename)
// 	ext := strings.Split(fileHeader.Filename, ".")
// 	src, err := fileHeader.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	// Create an absolute path for saving the file
// 	destinationDir := "/my-auth-app/destination" // Replace with the absolute path to your 'destination' directory
// 	dstPath := filepath.Join(destinationDir, ext[0]+strconv.Itoa(key)+"."+ext[1])

// 	dst, err := os.Create(dstPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

//		fmt.Printf("File saved to: %s\n", dstPath)
//		return nil
//	}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", handleUpload)
	r.GET("/image-list", listImages)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

func handleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the form data to JSON and print it
	formData := make(map[string]interface{})
	formData["files"] = form.File
	fmt.Printf("Form Data: %+v\n", formData)

	files := form.File["files"]
	for key, file := range files {
		err := saveFile(file, key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}

func saveFile(fileHeader *multipart.FileHeader, key int) error {
	fmt.Println(fileHeader.Filename)
	ext := strings.Split(fileHeader.Filename, ".")
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Get the current working directory and create an absolute path for saving the file
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	destinationDir := filepath.Join(currentDir, "destination")

	// Ensure the destination directory exists, create it if it doesn't
	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
			return err
		}
	}

	dstPath := filepath.Join(destinationDir, ext[0]+strconv.Itoa(key)+"."+ext[1])

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	fmt.Printf("File saved to: %s\n", dstPath)
	return nil
}
func listImages(c *gin.Context) {
	// Get a list of all image files in the "destination" directory
	destinationDir := "destination"
	files, err := ioutil.ReadDir(destinationDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Collect the list of image filenames
	var imageList []string
	for _, file := range files {
		// Check if the file is an image (you can customize this check)
		if strings.HasSuffix(file.Name(), ".png") || strings.HasSuffix(file.Name(), ".jpg") {
			imageList = append(imageList, file.Name())
		}
	}

	// Return the list of image filenames as a JSON response
	c.JSON(http.StatusOK, gin.H{"imageList": imageList})
}
