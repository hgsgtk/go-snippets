package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// custom template function
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	// default
	// r := gin.Default()

	// logging to a file
	// TODO: logging structure
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// using middleware
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// custom template function
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// revovery middleware recovery panic
	// curl http://localhost:8080/panic -v
	r.GET("/panic", func(c *gin.Context) {
		panic("if panic occured, return 500.")
	})
	// parameters in path
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// querystring parameters
	// curl http://localhost:8080/welcome?firstname=Kazuki&lastname=Higashiguchi
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// multipart/urlencoded form
	// curl -X POST -F 'message=hello' -F 'nick=kazuki' http://localhost:8080/form_post
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	// querystring paramters and post form
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(200, gin.H{
			"status":  "posted",
			"id":      id,
			"name":    name,
			"message": message,
			"page":    page,
		})
	})
	// upload single file
	// curl -X POST http://localhost:8080/upload -F "file=@/Users/users/Downloads/kurapika.png" -H "Content-Type: multipart/form-data"
	r.MaxMultipartMemory = 8 << 20
	r.Static("/file", "./public")
	r.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s", file.Filename, name, email))
	})
	// grouping routes
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong v1",
			})
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong v2",
			})
		})
	}
	// TODO: authorized group
	// authorized := r.Group("/admin")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.GET("/ping", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "ping authorized",
	// 		})
	// 	})
	// }
	// binding JSON
	// curl -v -X POST http://localhost:8080/loginJSON -H 'content-type: application/json' -d '{"user": "menu", "password": 123}'
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "menu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}
		}
	})
	// binding HTML form
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "menu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}
		}
	})
	// html template rendering
	r.LoadHTMLGlob("templates/**/*")
	// single directory
	// r.GET("/web", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })
	// multiple directory
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	// custom template function
	r.GET("/records/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "records/index.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
		// c.HTML(http.StatusOK, "times/index.tmpl", gin.H{
		// 	"now": ,
		// })
	})
	r.Run(":8080")
}
