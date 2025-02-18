package mocks

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"gorm.io/gorm"
)

func GenerateListOpenings(quantity int) []schemas.Opening {
	var mockListOpenings []schemas.Opening
	rand.Seed(time.Now().UnixNano())

	languages := []string{"Java", "GO", "JavaScript", "Flutter", "React"}
	companies := []string{"Tech Corp", "+3000 DEV", "Future Tech", "You in Control"}
	countries := []string{"Spain", "USA", "Canada", "Portugal", "Germany"}

	for i := 1; i <= quantity; i++ {
		mockListOpenings = append(mockListOpenings, schemas.Opening{
			Model:    gorm.Model{ID: uint(i)},
			Role:     languages[rand.Intn(len(languages))] + " Developer",
			Company:  companies[rand.Intn(len(companies))],
			Location: countries[rand.Intn(len(countries))],
			Link:     "https://joblisting.com/job" + fmt.Sprint(i),
			Remote:   rand.Intn(2) == 1,
			Salary:   int64(rand.Intn(50000) + 50000),
		})
	}
	return mockListOpenings
}
