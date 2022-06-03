package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wilmerpino/mutant/domain/model"
)

type MockMutantRepository struct {
	mock.Mock
}

func (m *MockMutantRepository) CountAll() []model.StatsResult {
	args := m.Called()
	return args.Get(0).([]model.StatsResult)
}

func (m *MockMutantRepository) Save(info model.DnaInfo) error {
	args := m.Called(info)
	return args.Error(0)
}
