package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/usecase/interactor"
)

type healthCheckController struct {
	healthCheckInteractor interactor.IHealthCheckInteractor
}

type IHealthCheckController interface {
	GetHealthCheck(iris.Context)
}

func NewHealthCheckController(hci interactor.IHealthCheckInteractor) IHealthCheckController {
	return &healthCheckController{hci}
}

func (hc *healthCheckController) GetHealthCheck(ctx iris.Context) {
	obj, err := hc.healthCheckInteractor.GetHealthCheck()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.StopExecution()
		return
	}

	ctx.JSON(obj)
}
