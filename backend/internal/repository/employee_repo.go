package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) domain.EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) FindAll() ([]domain.Employee, error) {
	var employees []domain.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) FindByID(id uint) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.First(&employee, id).Error
	return &employee, err
}

func (r *employeeRepository) Create(employee *domain.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) Update(employee *domain.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Employee{}, id).Error
}
