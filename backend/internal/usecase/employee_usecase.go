package usecase

import "backend/internal/domain"

type employeeUsecase struct {
	repo domain.EmployeeRepository
}

func NewEmployeeUsecase(repo domain.EmployeeRepository) domain.EmployeeUsecase {
	return &employeeUsecase{repo}
}

func (u *employeeUsecase) GetAll() ([]domain.Employee, error) {
	return u.repo.FindAll()
}

func (u *employeeUsecase) GetByID(id uint) (*domain.Employee, error) {
	return u.repo.FindByID(id)
}

func (u *employeeUsecase) Create(employee *domain.Employee) error {
	return u.repo.Create(employee)
}

func (u *employeeUsecase) Update(employee *domain.Employee) error {
	return u.repo.Update(employee)
}

func (u *employeeUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
