package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryohei1216/firebase-learn/api"
	"github.com/ryohei1216/firebase-learn/repository"
	"github.com/ryohei1216/firebase-learn/service"
	"github.com/ryohei1216/firebase-learn/usecase"
)

func main() {
	r := gin.Default()

	fc := api.NewFirebaseClient()
	userRepository := repository.NewUserRepository(fc)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userService := service.NewUserService(userUsecase)

	r.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		u, err := userService.GetUser(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(200, gin.H{
			"user": u,
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}