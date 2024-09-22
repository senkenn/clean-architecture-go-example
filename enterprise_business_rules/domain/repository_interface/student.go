package repository_interface

import "clean-architecture-go-example/enterprise_business_rules/domain/entity"

type IStudentRepository interface {
	GetByID(id int) (*entity.Student, error)
	Save(student *entity.Student) error
}
