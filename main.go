package main

import (
	"github.com/denizcamalan/movie_app/docs"
	"github.com/denizcamalan/movie_app/router"
	log "github.com/siruspen/logrus"
	swgFiles "github.com/swaggo/files"
	swgGin "github.com/swaggo/gin-swagger"
)

var port = "8080"

type App struct {
	Name    string
	Version string
}

// Run - sets up our application and starts the server.
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		},
	).Info("Setting up application")

	return nil
}

func main() {
	app := &App{
		Name:    "movie_app",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Failed to run application")
		log.Fatal(err)
	}

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
	// @description JWT Authorization header using the Bearer scheme. Example: \"Bearer {token}\"

	router := router.NewRoutes()

	url := swgGin.URL("http://localhost:" + port + "/swagger/doc.json")
	router.GET("/swagger/*any", swgGin.WrapHandler(swgFiles.Handler, url))

	router.Run(":" + port)

}