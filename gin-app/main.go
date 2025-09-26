package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// simple
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// path parameter
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})

	// query parameter
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		c.String(http.StatusOK, "Search for %s", query)
	})

	// form parameter
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// JSON parameter
	r.POST("/register", func(c *gin.Context) {
		var user struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"reqiured,email"`
			Password string `json:"password" binding:"required,min=6"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	})

	// book group api
	book := r.Group("/book")
	{
		book.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name":  "Golang",
				"price": 320.00,
			})
		})

		book.POST("/", func(c *gin.Context) {
			var book struct {
				Name  string  `json:"name"`
				Price float32 `json:"price"`
			}

			if err := c.ShouldBindJSON(&book); err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "book added",
			})
		})
	}

	// with middleware
	course := r.Group("/course").Use(AuthMiddleware())
	{
		course.GET("/", func(c *gin.Context) {
			var course struct {
				Name  string  `json:"name"`
				Price float32 `json:"price"`
			}

			if err := c.ShouldBindJSON(&course); err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "course added",
			})
		})
	}

	// file upload
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("image")
		c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"message": "File Uploaded successfully!" + file.Filename,
		})
	})

	// load html template
	r.LoadHTMLGlob("templates/*")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "Welcome to Gin Templates",
			"name":  "Uttam Nath",
			"items": []string{"Go", "Gin", "GORM"},
		})
	})

	r.Run(":8080")

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "test12345" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
