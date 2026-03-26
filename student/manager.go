package student

import "errors"

type Manager struct {
	students []Student
}

func NewManager() *Manager {
	return &Manager{
		students: []Student{},
	}
}

func (m *Manager) AddStudent(s Student) error {
	// 🔍 檢查是否已存在相同 ID
	for _, stu := range m.students {
		if stu.ID == s.ID {
			return errors.New("ID已存在，不能重複新增")
		}
	}

	m.students = append(m.students, s)
	return nil
}

func (m *Manager) GetStudentById(id string) (Student, error) {
	for _, s := range m.students {
		if s.ID == id {
			return s, nil
		}
	}
	return Student{}, errors.New("找不到學生")
}

func (m *Manager) GetAll() []Student {
	return m.students
}

func (m *Manager) UpdateStudent(id string, newData Student) error {
	for i, s := range m.students {
		if s.ID == id {
			m.students[i] = newData
			return nil
		}
	}
	return errors.New("找不到學生")
}

func (m *Manager) DeleteStudent(id string) error {
	for i, s := range m.students {
		if s.ID == id {
			m.students = append(m.students[:i], m.students[i+1:]...)
			return nil
		}
	}
	return errors.New("找不到學生")
}
