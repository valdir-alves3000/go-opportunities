package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
)

type OpeningRepositoryMock struct {
	mock.Mock
}

func (m *OpeningRepositoryMock) Create(opening schemas.Opening) error {
	args := m.Called(opening)
	return args.Error(0)
}

func (m *OpeningRepositoryMock) FindByID(id uint) (*schemas.Opening, error) {
	args := m.Called(id)
	return args.Get(0).(*schemas.Opening), args.Error(1)
}

func (m *OpeningRepositoryMock) Update(opening schemas.Opening) error {
	args := m.Called(opening)
	return args.Error(0)
}

func (m *OpeningRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *OpeningRepositoryMock) FindAll(limit, offset int) ([]schemas.Opening, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]schemas.Opening), args.Error(1)
}
