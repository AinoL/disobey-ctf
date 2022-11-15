package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"crypto/md5"
	"path/filepath"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	fileDirectory := "./images"

	r.POST("image", func(c *gin.Context) {
		message := c.PostForm("message")
		resp, err := http.Get(message)
		if err != nil {
			print(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			print(err)
		}
		
		hash := fmt.Sprintf("%x", md5.Sum(body))
		print(hash)

		fileName := filepath.Join(fileDirectory, hash)

		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			print(err)
		}
		_, err = file.Write(body)
		if err != nil {
			print(err)
		}
		c.String(http.StatusOK, string(fileName))
	})

	r.Static("/images", fileDirectory)
	r.StaticFile("/", "./web/index.html")

	return r
}
