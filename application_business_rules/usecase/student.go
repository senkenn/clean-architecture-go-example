package usecase

import (
	"clean-architecture-go-example/enterprise_business_rules/domain"
	"clean-architecture-go-example/enterprise_business_rules/domain/repository_interface"
)

type StudentUseCase struct {
	StudentRepo repository_interface.IStudentRepository
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Get student by ID
func (uc *StudentUseCase) GetStudentByID(id int) (*domain.Student, error) {
	student, err := uc.StudentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

// Create a new student
func (uc *StudentUseCase) CreateStudent(id int, name string) error {
	student := &domain.Student{
		ID:   id,
		Name: name,
	}
	return uc.StudentRepo.Save(student)
}
