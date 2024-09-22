package repository

import (
	"clean-architecture-go-example/enterprise_business_rules/domain"
	"errors"
	"fmt"
)

// This is just an in-memory mock, you can replace it with real DB implementation
type InMemoryStudentRepository struct {
	students map[int]*domain.Student
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		students: make(map[int]*domain.Student),
	}
}

func (repo *InMemoryStudentRepository) GetByID(id int) (*domain.Student, error) {
	if student, exists := repo.students[id]; exists {
		return student, nil
	}
	return nil, errors.New("student not found")
}

func (repo *InMemoryStudentRepository) Save(student *domain.Student) error {
	repo.students[student.ID] = student
	fmt.Println("Student saved:", student)
	return nil
}
