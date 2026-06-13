package domain

import "time"

type Employee struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Position  string    `json:"position"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmployeeRepository interface {
	FindAll() ([]Employee, error)
	FindByID(id uint) (*Employee, error)
	Create(employee *Employee) error
	Update(employee *Employee) error
	Delete(id uint) error
}

type EmployeeUsecase interface {
	GetAll() ([]Employee, error)
	GetByID(id uint) (*Employee, error)
	Create(employee *Employee) error
	Update(employee *Employee) error
	Delete(id uint) error
}
