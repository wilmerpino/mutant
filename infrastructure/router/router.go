package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/wilmerpino/mutant/interface/controller"
)

func NewRouter(app *iris.Application, h controller.HealthCheckController, c controller.AppController) {
	app.Use(logger.New())
	app.Use(recover.New())

	// Configure healthcheck
	app.Get("/healthcheck", h.GetHealthCheck)
	app.Post("/mutant", c.IsMutant)
	app.Get("/stats", c.GetStats)
}
