package framework

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type App struct {
	container *dig.Container
	router    *gin.Engine
}

func New(modules ...Module) *App {
	app := &App{
		container: dig.New(),
		router:    gin.Default(),
	}

	for _, mod := range modules {
		if err := mod.Register(app.container); err != nil {
			log.Fatalf("Grob: Failed to register module: %v", err)
		}
	}

	log.Println("Grob App Intitialized")
	return app
}

func (a *App) Start() {
	log.Println("Grob: Starting server on http://localhost:8080")
	a.router.Run(":8080")
}
