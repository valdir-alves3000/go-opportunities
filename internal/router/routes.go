package router

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/valdir-alves3000/go-opportunities/config"
	docs "github.com/valdir-alves3000/go-opportunities/docs"
	"github.com/valdir-alves3000/go-opportunities/internal/core/usecases/opening_usecase"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/repositories"
)

const BASE_PATH = "/api/v1"

func initializeRoutes(r *gin.Engine) {
	db := config.GetSQLite()
	opRepo := repositories.NewOpeningRepository(db)
	opUsecase := opening_usecase.NewOpeningUseCase(opRepo)
	opHandler := handler.NewOpeningHandler(opUsecase)

	v1 := r.Group(BASE_PATH)
	{
		v1.GET("/openings", opHandler.List)
		v1.GET("/openings/:id", opHandler.ShowOpening)
		v1.POST("/openings", opHandler.Create)
		v1.DELETE("/openings/:id", opHandler.Delete)
		v1.PUT("/openings/:id", opHandler.Update)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
}

func setupSwagger(r *gin.Engine) {
	swaggerHost := getSwaggerHost()
	docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.BasePath = BASE_PATH

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func getSwaggerHost() string {
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "localhost:8080"
	}
	return host
}
