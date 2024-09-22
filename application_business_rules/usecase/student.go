package usecase

import "clean-architecture-go-example/enterprise_business_rules/domain"

type StudentUseCase struct {
	StudentRepo domain.StudentRepository
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
