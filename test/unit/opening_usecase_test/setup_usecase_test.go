package opening_usecase_test

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/usecases/opening_usecase"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
)

func boolPtr(b bool) *bool {
	return &b
}

func setupUsecaseTest() (*opening_usecase.OpeningUseCase, *mocks.OpeningRepositoryMock) {
	repo := new(mocks.OpeningRepositoryMock)
	uc := opening_usecase.NewOpeningUseCase(repo)

	return uc, repo
}
