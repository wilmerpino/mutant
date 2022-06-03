package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wilmerpino/mutant/domain/model"
)

type MockHealthCheckInteractor struct {
	mock.Mock
}

func (m *MockHealthCheckInteractor) GetHealthCheck() (model.HealthCheckResponse, error) {
	args := m.Called()
	return args.Get(0).(model.HealthCheckResponse), args.Error(1)
}

type MockMutantInteractor struct {
	mock.Mock
}

func (m *MockMutantInteractor) Stats() (model.Stats, error) {
	args := m.Called()
	return args.Get(0).(model.Stats), args.Error(1)
}

func (m *MockMutantInteractor) SaveStrand(data model.DnaInfo) error {
	args := m.Called(data)
	return args.Error(0)
}
