package opening_usecase

import (
	"fmt"

	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
)

func errParamIsRequired(name, typ string) *internal_error.InternalError {
	message := fmt.Sprintf("param: %s (type: %s) is required", name, typ)
	return internal_error.NewBadRequestError(message)
}

func (uc *OpeningUseCase) Create(co schemas.CreateOpeningRequest) *internal_error.InternalError {
	err := validate(&co)
	if err != nil {
		return err
	}

	opening := schemas.Opening{
		Role:     co.Role,
		Company:  co.Company,
		Location: co.Location,
		Remote:   *co.Remote,
		Link:     co.Link,
		Salary:   co.Salary,
	}

	errRepo := uc.repo.Create(opening)
	if errRepo != nil {
		return internal_error.NewInternalServerError("error creating opening")
	}

	return nil
}

func validate(co *schemas.CreateOpeningRequest) *internal_error.InternalError {
	requiredFields := map[string]interface{}{
		"role":     co.Role,
		"company":  co.Company,
		"location": co.Location,
		"link":     co.Link,
	}

	for field, value := range requiredFields {
		if value == "" {
			return errParamIsRequired(field, "string")
		}
	}

	if co.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if co.Salary < 3000 {
		message := "salary must be at least 3k"
		return internal_error.NewBadRequestError(message)
	}

	return nil
}
