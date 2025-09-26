package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int64  `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Age   byte   `json:"age" binding:"required"`
	Phone string `json:"phone"`
}

var users = []User{
	{Id: 1, Name: "Uttam Nath", Age: 24, Phone: "9087654321"},
	{Id: 2, Name: "Prince Nath", Age: 19},
}

func main() {

	r := gin.Default()

	// get all users
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// get user by id
	r.GET("/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID!",
			})
		}

		for _, user := range users {
			if user.Id == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
	})

	// add user
	r.POST("/", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{
			"message": "User added!",
		})
	})

	// update user
	r.PUT("/:id", func(c *gin.Context) {
		var user User
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID!",
			})
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		for i, u := range users {
			if u.Id == id {
				users[i].Name = user.Name
				users[i].Age = user.Age
				users[i].Phone = user.Phone
				c.JSON(http.StatusOK, gin.H{
					"message": "User updated!",
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
	})

	// delete user
	r.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID!",
			})
		}

		for i, user := range users {
			if user.Id == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message": "User deleted!",
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not find!",
		})
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("error: " + err.Error())
	}
}
