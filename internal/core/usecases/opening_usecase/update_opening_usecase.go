package opening_usecase

import (
	"reflect"

	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"gorm.io/gorm"
)

func (uc *OpeningUseCase) Update(id uint, upo schemas.UpdateOpeningRequest) *internal_error.InternalError {
	err := validateUpdateOpeningRequest(&upo)
	if err != nil {
		return err
	}

	if err := validateSalary(upo.Salary); err != nil {
		return err
	}

	opening, errRepo := uc.repo.FindByID(id)
	if errRepo != nil {
		return internal_error.NewNotFoundError("opening not found")
	}

	upOpening := schemas.Opening{
		Model:    gorm.Model{ID: id},
		Role:     getFieldValue(upo.Role, opening.Role),
		Company:  getFieldValue(upo.Company, opening.Company),
		Location: getFieldValue(upo.Location, opening.Location),
		Link:     getFieldValue(upo.Link, opening.Link),
		Remote:   getRemoteValue(upo.Remote, opening.Remote),
		Salary:   upo.Salary,
	}

	errRepo = uc.repo.Update(upOpening)
	if errRepo != nil {
		return internal_error.NewInternalServerError("error updating opening")
	}
	return nil
}

func validateUpdateOpeningRequest(upo *schemas.UpdateOpeningRequest) *internal_error.InternalError {
	v := reflect.ValueOf(*upo)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() != "" {
			return nil
		}
	}

	if upo.Remote != nil {
		return nil
	}

	if upo.Salary != 0 {
		return nil 
	}

	return internal_error.NewBadRequestError("at least one valid field must be provided")
}

func getFieldValue(updated, original string) string {
	if updated != "" {
		return updated
	}
	return original
}

func getRemoteValue(updated *bool, original bool) bool {
	if updated != nil {
		return *updated
	}
	return original
}

func validateSalary(salary int64) *internal_error.InternalError {
	if salary < 3000 {
		return internal_error.NewBadRequestError("salary must be at least 3k")
	}
	return nil
}
