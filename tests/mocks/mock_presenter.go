package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wilmerpino/mutant/domain/model"
)

type MockHealthCheckPresenter struct {
	mock.Mock
}

func (m *MockHealthCheckPresenter) ResponseHealthCheck() model.HealthCheckResponse {
	args := m.Called()
	return args.Get(0).(model.HealthCheckResponse)
}

type MockMutantPresenter struct {
	mock.Mock
}

func (m *MockMutantPresenter) ResponseMutantsStats(ac []model.StatsResult) model.Stats {
	args := m.Called(ac)
	return args.Get(0).(model.Stats)
}
