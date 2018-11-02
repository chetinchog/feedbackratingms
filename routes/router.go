package routes

import (
	"github.com/chetinchog/feedbackratingms/controllers"
	"github.com/gin-gonic/gin"
)

func Route(server *gin.Engine) {
	// System
	server.GET("/v1", controllers.Check)

	// Rules
	server.POST("/v1/rates/:articleId/rules", controllers.SetRules)
	server.GET("/v1/rates/:articleId/rules", controllers.GetRules)

	// Rates
	server.GET("/v1/rates/:articleId/", controllers.GetRate)
	server.GET("/v1/rates/:articleId/history", controllers.GetHistory)
}
