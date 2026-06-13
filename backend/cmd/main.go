package main

import (
	"log"

	"backend/config"
	deliveryHttp "backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&domain.Employee{})

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeRepo)
	router := deliveryHttp.NewRouter(employeeUsecase)

	log.Printf("server running on port %s", cfg.AppPort)
	router.Run(":" + cfg.AppPort)
}
