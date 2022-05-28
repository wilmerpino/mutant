package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IHealthCheckPresenter interface {
	ResponseHealthCheck() model.HealthCheckResponse
}
