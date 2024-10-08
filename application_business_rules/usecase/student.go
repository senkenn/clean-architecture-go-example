package usecase

import (
	"clean-architecture-go-example/enterprise_business_rules/domain/entity"
	"clean-architecture-go-example/enterprise_business_rules/domain/repository_interface"
)

type StudentUseCase struct {
	StudentRepo repository_interface.IStudentRepository
}

// Get student by ID
func (uc *StudentUseCase) GetStudentByID(id int) (*entity.Student, error) {
	student, err := uc.StudentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

// Create a new student
func (uc *StudentUseCase) CreateStudent(id int, name string) error {
	student := &entity.Student{
		ID:   id,
		Name: name,
	}
	return uc.StudentRepo.Save(student)
}
