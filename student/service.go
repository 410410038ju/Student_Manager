package student

type StudentService interface {
	AddStudent(s Student) error
	GetStudentById(id string) (Student, error)
	GetAll() []Student
	UpdateStudent(id string, s Student) error
	DeleteStudent(id string) error
}
