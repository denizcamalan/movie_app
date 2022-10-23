package main

import (
	"github.com/denizcamalan/movie_app/docs"
	"github.com/denizcamalan/movie_app/router"
	logs "github.com/sirupsen/logrus"
	swgFiles "github.com/swaggo/files"
	swgGin "github.com/swaggo/gin-swagger"
)


type App struct {
	Name    string
	Version string
}

// Run - sets up our application and starts the server.
func (a *App) Run() error {
	logs.SetFormatter(&logs.JSONFormatter{})
	logs.WithFields(
		logs.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		},
	).Info("Setting up application")
	
	return nil
}

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

	app := &App{
		Name:    "movie_app",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		logs.Error("Failed to run application")
		logs.Fatal(err)
	}
}
