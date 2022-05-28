package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/usecase/interactor"
)

type healtCheckController struct {
	healtCheckInteractor interactor.IHealtCheckInteractor
}

type IHealtCheckController interface {
	GetHealthCheck(iris.Context)
}

func NewHealtCheckController(hci interactor.IHealtCheckInteractor) IHealtCheckController {
	return &healtCheckController{hci}
}

func (hc *healtCheckController) GetHealthCheck(ctx iris.Context) {
	obj, err := hc.healtCheckInteractor.GetHealthCheck()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.StopExecution()
		return
	}

	ctx.JSON(obj)
}
