package main

import (
	"github.com/denizcamalan/movie_app/docs"
	"github.com/denizcamalan/movie_app/router"
	swgFiles "github.com/swaggo/files"
	swgGin "github.com/swaggo/gin-swagger"
)


func main() {

	port := "8080"

	// setup swagger documantation

	docs.SwaggerInfo.Title = "Swagger Service Movie App"
	docs.SwaggerInfo.Description = "This is service API documentation."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/movie_archive"
	docs.SwaggerInfo.Schemes = []string{"http"}
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	router := router.NewRoutes()
	url := swgGin.URL("http://localhost:" + port + "/swagger/doc.json")
	router.GET("/swagger/*any", swgGin.WrapHandler(swgFiles.Handler, url))
	router.Run(":" + port)

}
