package routes

import (
	"github.com/DHRUVIT18/go-jwt-project/controllers"
	"github.com/DHRUVIT18/go-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())

}
