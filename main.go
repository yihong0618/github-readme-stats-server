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
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
	r.GET("/:username", func(c *gin.Context) {
		r.LoadHTMLGlob("templates/*.html")
		NameList := makeUserNameList()
		name := c.Param("username")
		if ContainsInArray(strings.ToLower(name), NameList) {
			c.HTML(http.StatusOK, name+".html", nil)
		} else {
			c.HTML(http.StatusOK, "homepage.html", nil)
		}
	})
	r.POST("/generate", func(c *gin.Context) {
		needRefresh := false
		userName, _ := c.GetPostForm("r")
		userName = strings.TrimSpace(userName)
		l := strings.Split(userName, "--")
		if len(l) > 1 && l[1] == "refresh" {
			needRefresh = true
		}
		userName = l[0]
		userName = strings.TrimSpace(userName)
		// TODO refactor
		userNameList := makeUserNameList()
		if ContainsInArray(userName, userNameList) && !needRefresh {
			userName = strings.ToLower(userName)
			c.HTML(http.StatusOK, userName+".html", nil)
		} else {
			result := github.GenerateNewFile(userName)
			// warit for a while to make sure the file is generated
			time.Sleep(time.Second * 2)
			userName = strings.ToLower(userName)
			// c.HTML(http.StatusOK, userName+".html", nil)
			c.Data(http.StatusOK, "text/html; charset=utf-8", result)
		}
	})
	r.Run()
}
