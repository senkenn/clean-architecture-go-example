package repository_interface

import "clean-architecture-go-example/enterprise_business_rules/domain"

type IStudentRepository interface {
	GetByID(id int) (*domain.Student, error)
	Save(student *domain.Student) error
}
