package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
)

type OpeningUseCaseMock struct {
	mock.Mock
}

func (m *OpeningUseCaseMock) Create(co schemas.CreateOpeningRequest) *internal_error.InternalError {
	args := m.Called(co)
	return args.Get(0).(*internal_error.InternalError)
}

func (m *OpeningUseCaseMock) DeleteByID(id uint) *internal_error.InternalError {
	args := m.Called(id)
	return args.Get(0).(*internal_error.InternalError)
}

func (m *OpeningUseCaseMock) GetByID(id uint) (*schemas.Opening, *internal_error.InternalError) {
	args := m.Called(id)
	return args.Get(0).(*schemas.Opening), args.Get(1).(*internal_error.InternalError)
}

func (m *OpeningUseCaseMock) Update(id uint, upo schemas.UpdateOpeningRequest) *internal_error.InternalError {
	args := m.Called(id)
	return args.Get(0).(*internal_error.InternalError)
}

func (m *OpeningUseCaseMock) ListOpenings(page int) ([]schemas.Opening, *internal_error.InternalError) {
	args := m.Called(page)
	return args.Get(0).([]schemas.Opening), args.Get(1).(*internal_error.InternalError)
}
