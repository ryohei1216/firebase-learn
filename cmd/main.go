package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryohei1216/firebase-learn/api"
	"github.com/ryohei1216/firebase-learn/infrastructure"
	"github.com/ryohei1216/firebase-learn/service"
	"github.com/ryohei1216/firebase-learn/usecase"
)

func main() {
	r := gin.Default()

	fc := api.NewAuthClient()
	sc := api.NewStoreClient()

	// TODO: create時の対策考える
	// r.Use(api.AuthMiddleware(fc))

	userRecordRepository := infrastructure.NewUserRecordRepository(fc)
	userRepository := infrastructure.NewUserRepository(sc)

	userUsecase := usecase.NewUserUsecase(userRecordRepository, userRepository)
	userService := service.NewUserService(userUsecase)

	r.POST("/users", func(c *gin.Context) {
		var json struct{
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := userService.CreateUser(c.Request.Context(), json.Email, json.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": u})
	})

	r.GET("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		u, err := userService.GetUser(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	})

	r.POST("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		var json struct{
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := userService.UpdateUser(c.Request.Context(), uid, json.Email, json.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": u})
	})

	r.DELETE("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		err := userService.DeleteUser(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": "deleted"})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
