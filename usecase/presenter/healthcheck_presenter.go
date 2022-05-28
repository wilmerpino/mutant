package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IHealtCheckPresenter interface {
	ResponseHealthCheck() model.HealthCheckResponse
}
