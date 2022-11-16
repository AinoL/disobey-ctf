package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	fileDirectory := "./images"
	staticDirectory := "./static"

	images := []string{
		"/images/35bc566eb29102e880ae4f231458ca66",
		"/images/35bc566eb29102e880ae4f231458ca66",
		"/images/35bc566eb29102e880ae4f231458ca66",
		"/images/35bc566eb29102e880ae4f231458ca66",
		"/images/35bc566eb29102e880ae4f231458ca66",
		"/images/35bc566eb29102e880ae4f231458ca66",
	}

	r.POST("image", func(c *gin.Context) {
		message := c.PostForm("url")
		resp, err := http.Get(message)
		if err != nil {
			c.String(http.StatusBadRequest, "err")
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.String(http.StatusBadRequest, "err")
			return
		}

		hash := fmt.Sprintf("%x", md5.Sum(body))
		fileName := filepath.Join(fileDirectory, hash)

		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			c.String(http.StatusBadRequest, "err")
			return
		}
		_, err = file.Write(body)
		if err != nil {
			c.String(http.StatusBadRequest, "err")
			return
		}
		host := c.Request.Host
		url := fmt.Sprintf("http://%s/images/%s", host, hash)

		images = append(images, fmt.Sprintf("/images/%s", hash))

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Url":    url,
			"Images": images,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Images": images,
		})
	})

	r.Static("/images", fileDirectory)
	r.Static("/static", staticDirectory)

	return r
}
