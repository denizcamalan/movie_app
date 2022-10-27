package router

import (
	"github.com/denizcamalan/movie_app/controller"
	"github.com/denizcamalan/movie_app/middleware"
	"github.com/gin-gonic/gin"
)

var api_controller = &controller.API{}

func NewRoutes() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/movie_archive")

	v1.GET("/ping", Ping)
	
	RoutesPost(v1)
	return router
}

func RoutesPost(rg *gin.RouterGroup) {
	rg.GET("/list",api_controller.GetAll)
	rg.POST("/register", api_controller.Register)
	rg.POST("/login",api_controller.Login)

	secured := rg.Group("/admin").Use(middleware.JwtAuthMiddleware())
	{
		secured.GET("/user", api_controller.CurrentUser)
		secured.POST("/add", api_controller.AddData)
		secured.PUT("/update/:id", api_controller.UpdateData)
		secured.GET("/get/:id", api_controller.GetDataById)
		secured.DELETE("/delete/:id",api_controller.DeleteDataByID)
	}
	
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}