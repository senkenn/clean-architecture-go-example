package model

import "clean-architecture-go-example/enterprise_business_rules/domain/entity"

type Student struct {
	entity.Student
}

func (su *Student) IsValid() bool {
	return true
}
