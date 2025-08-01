package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"watermark-image/aws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

const MSG_ERROR_DELETE_FILE = "Unable to delete file\n %q, %v"

func main() {
	godotenv.Load()

	router := gin.New()

	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXIES")})

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGINS")},
		AllowMethods: []string{"POST"},
		MaxAge:       12 * time.Hour,
	}))

	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		filename := file.Filename

		c.SaveUploadedFile(file, os.Getenv("TMP_FOLDER")+filename)

		errorUpload := aws.Upload(filename)

		if errorUpload == nil {
			err := os.Remove(os.Getenv("TMP_FOLDER") + filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, MSG_ERROR_DELETE_FILE, filename, err)
			}
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", filename))
		} else {
			c.String(http.StatusBadRequest, fmt.Sprintf("'%s'", errorUpload.Error()))
		}
	})

	router.Run(os.Getenv("PORT"))
}
