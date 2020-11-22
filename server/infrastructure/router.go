package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindai-csg/D-Chat/interfaces/controller"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	mongoHandler, err := NewMongoHandler()
	if err != nil {
		panic(err)
	}
	logger := NewLoggerStd()
	jwt := NewTokenJwt("secret", 24)

	userController := controller.NewUserController(logger, mongoHandler)
	messageController := controller.NewMessageController(logger, mongoHandler)
	tokenController := controller.NewTokenController(logger, jwt, mongoHandler)

	// ------------------------------
	// api v1
	// ------------------------------
	v1 := router.Group("/v1")
	// parameter: gateway.LoginInput
	// response: gateway.TokenOutput
	v1.POST("/login", func(c *gin.Context) { tokenController.CreateToken(c) })
	// parameter: gateway.UserInput
	// response: gateway.UserInfoOutput
	v1.POST("/users", func(c *gin.Context) { userController.CreateUser(c) })
	// parameter: gateway.MessageInput
	// response: gateway.StatusMessageOutput
	v1.POST("/messages", func(c *gin.Context) { messageController.CreateMessage(c) })

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	Router = router
}
