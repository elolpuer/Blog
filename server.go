package main

import (
	"log"

	"github.com/elolpuer/blog/pkg/controller"
	authController "github.com/elolpuer/blog/pkg/controller/auth_controller"
	postController "github.com/elolpuer/blog/pkg/controller/post_controller"
	userController "github.com/elolpuer/blog/pkg/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Index controller
	router.GET("/", controller.IndexGet)
	//Post controller
	router.GET("/posts", postController.PostsGet)
	router.POST("/add/process", postController.AddPost)
	router.POST("/delete", postController.DeletePost)
	//Auth controller
	router.GET("/signup", authController.SignUpGet)
	router.GET("/signin", authController.SignInGet)
	router.POST("/signup/auth", authController.SignUpPost)
	router.POST("/signin/auth", authController.SignInPost)
	router.POST("/logout", authController.LogOutPost)
	//User controller
	router.GET("/user", userController.UserGet)
	router.GET("/users", userController.UsersGet)

	err := router.Run(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
