package main

import (
	"fmt"

	"github.com/iris-contrib/swagger/v12"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/v12/swaggerFiles" // swagger embed files
	"github.com/kataras/iris/v12"
	_ "github.com/wilmerpino/mutant/docs"
	"github.com/wilmerpino/mutant/infrastructure/config"
	"github.com/wilmerpino/mutant/infrastructure/database"
	"github.com/wilmerpino/mutant/infrastructure/router"
	"github.com/wilmerpino/mutant/registry"
)

func main() {
	v := config.EnviromentConfig{}
	v.InitVariables()

	db := database.NewDB(v.Database)

	app := iris.Default()
	r := registry.NewRegistry(db)

	router.NewRouter(app, r.NewHealthController(), r.NewAppController())

	// Swagger Doc
	swaggerConfig := &swagger.Config{
		URL: v.SwaggerHost + "/swagger/doc.json", // The url pointing to API definition
	}

	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

	host := ""
	if v.PortServer != "80" && v.PortServer != "443" {
		host = ":" + v.PortServer
	}
	if err := app.Listen(host); err != nil {
		app.Logger().Errorf("Unable to start server ")
	}

	app.Logger().Info(fmt.Sprintf(
		"Connected to the database %s in %s:%s",
		v.Database.Name,
		v.Database.Host,
		v.Database.Port,
	))

	app.Logger().Info(fmt.Sprintf("Server started in port %s", v.PortServer))
}
