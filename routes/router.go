package routes

import (
	"github.com/chetinchog/feedbackratingms/controllers"
	"github.com/gin-gonic/gin"
)

func Route(server *gin.Engine) {
	server.GET("/v1", controllers.Check)
}
