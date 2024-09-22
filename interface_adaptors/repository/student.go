package repository

import (
	"clean-architecture-go-example/enterprise_business_rules/domain/entity"
	"errors"
)

// This is just an in-memory mock, you can replace it with real DB implementation
type InMemoryStudentRepository struct {
	students map[int]*entity.Student
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		students: make(map[int]*entity.Student),
	}
}

func (repo *InMemoryStudentRepository) GetByID(id int) (*entity.Student, error) {
	if student, exists := repo.students[id]; exists {
		return student, nil
	}
	return nil, errors.New("student not found")
}

func (repo *InMemoryStudentRepository) Save(student *entity.Student) error {
	repo.students[student.ID] = student
	return nil
}
