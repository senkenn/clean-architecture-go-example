package domain

type StudentRepository interface {
	GetByID(id int) (*Student, error)
	Save(student *Student) error
}
