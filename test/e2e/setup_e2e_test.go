package e2e

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/core/usecases/opening_usecase"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	router   *gin.Engine
	db       *gorm.DB
	logger   *config.Logger
	basePath = "/api/v1"
)

func setupE2E() {
	logger = config.GetLogger("E2E")
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	dbPath := "./db/test.db"
	if _, err := os.Stat("./db"); os.IsNotExist(err) {
		err := os.MkdirAll("./db", 0755)
		if err != nil {
			panic(fmt.Sprintf("failed to create db directory: %v", err))
		}
	}

	file, err := os.Create(dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to create db file: %v", err))
	}
	file.Close()

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	opRepo := repositories.NewOpeningRepository(db)
	opUsecase := opening_usecase.NewOpeningUseCase(opRepo)
	opHandler := handler.NewOpeningHandler(opUsecase)

	// Route Definitions
	v1 := router.Group(basePath)
	{
		v1.POST("/openings", opHandler.Create)
		v1.GET("/openings", opHandler.List)
		v1.GET("/openings/:id", opHandler.ShowOpening)
		v1.DELETE("/openings/:id", opHandler.Delete)
		v1.PUT("/openings/:id", opHandler.Update)
	}
}

func tearDownE2E() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get database instance: %v", err))
	}

	sqlDB.Close()
	err = os.Remove("./db/test.db")
	if err != nil {
		logger.Errorf("Error removing test database: %v", err)
	}

	err = os.RemoveAll("./db")
	if err != nil {
		logger.Errorf("Error removing test database directory: %v", err)
	}
}

func TestMain(m *testing.M) {
	os.Setenv("CGO_ENABLED", "1")
	setupE2E()
	code := m.Run()
	tearDownE2E()

	os.Exit(code)
}
