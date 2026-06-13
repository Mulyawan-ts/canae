package http

import (
	"backend/internal/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(employeeUsecase domain.EmployeeUsecase) *gin.Engine {
	r := gin.Default()

	// cors — izinkan request dari frontend Astro
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4321"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// routes
	api := r.Group("/api/v1")
	{
		employee := NewEmployeeHandler(employeeUsecase)
		api.GET("/employees", employee.GetAll)
		api.GET("/employees/:id", employee.GetByID)
		api.POST("/employees", employee.Create)
		api.PUT("/employees/:id", employee.Update)
		api.DELETE("/employees/:id", employee.Delete)
	}

	return r
}
