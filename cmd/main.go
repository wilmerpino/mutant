package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
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

	if err := app.Listen(":" + v.PortServer); err != nil {
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
