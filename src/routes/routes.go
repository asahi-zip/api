package routes

import (
	"github.com/asahi-zip/api/controllers"
	"github.com/asahi-zip/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.TokenAuthMiddleware())
	protected.POST("/orgs/new", controllers.CreateOrg)
	protected.POST("/media/upload", controllers.UploadMedia)

	return r
}
