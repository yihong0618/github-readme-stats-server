package main

import (
	"fmt"
	"github/github"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var templatesDir = "templates"

func makeUserNameList() []string {
	fileList := []string{}
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		fileName := strings.Split(f.Name(), ".")
		fileList = append(fileList, fileName[0])
	}
	return fileList
}

func ContainsInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()
	userNameList := makeUserNameList()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "yihong0618.html", nil)
	})
	r.POST("/generate", func(c *gin.Context) {
		userName, _ := c.GetPostForm("r")
		if ContainsInArray(userName, userNameList) {
			c.HTML(http.StatusOK, userName+".html", nil)
		} else {
			github.GenerateNewFile(userName)
			// warit for a while to make sure the file is generated
			time.Sleep(time.Second * 1)
			c.HTML(http.StatusOK, userName+".html", nil)
		}
	})
	r.Run()
}
