package student

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Handler struct {
	Service StudentService
	Reader  *bufio.Reader
}

func NewHandler(service StudentService, reader *bufio.Reader) *Handler {
	return &Handler{
		Service: service,
		Reader:  reader,
	}
}

func (h *Handler) AddStudent() {
	fmt.Print("輸入ID: ")
	id, _ := h.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	// 🔍 先檢查學生是否存在
	_, er := h.Service.GetStudentById(id)
	if er == nil {
		fmt.Println("學生已存在")
		return
	}

	fmt.Print("輸入姓名: ")
	name, _ := h.Reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("輸入年齡: ")
	ageStr, _ := h.Reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, _ := strconv.Atoi(ageStr)

	fmt.Print("輸入班級: ")
	class, _ := h.Reader.ReadString('\n')
	class = strings.TrimSpace(class)

	student := Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Class: class,
	}

	err := h.Service.AddStudent(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("新增成功")
}

func (h *Handler) GetStudentById() {
	fmt.Print("輸入ID: ")
	id, _ := h.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	s, err := h.Service.GetStudentById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("學生ID: %s\n姓名: %s\n年齡: %d\n班級: %s\n", s.ID, s.Name, s.Age, s.Class)
}

func (h *Handler) GetAll() {
	students := h.Service.GetAll()

	if len(students) == 0 {
		fmt.Println("目前沒有學生資料")
		return
	}

	for _, s := range students {
		fmt.Printf("學生ID: %s\n姓名: %s\n年齡: %d\n班級: %s\n", s.ID, s.Name, s.Age, s.Class)
		fmt.Println("-----")
	}
}

func (h *Handler) UpdateStudent() {
	fmt.Print("輸入ID: ")
	id, _ := h.Reader.ReadString('\n')
	id = strings.TrimSpace(id)

	// 🔍 先檢查學生是否存在
	_, er := h.Service.GetStudentById(id)
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Print("輸入新姓名: ")
	name, _ := h.Reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("輸入新年齡: ")
	ageStr, _ := h.Reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, _ := strconv.Atoi(ageStr)

	fmt.Print("輸入新班級: ")
	class, _ := h.Reader.ReadString('\n')
	class = strings.TrimSpace(class)

	student := Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Class: class,
	}

	err := h.Service.UpdateStudent(id, student)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("更新成功")
}

func (h *Handler) DeleteStudent() {
	fmt.Print("輸入ID: ")
	id, _ := h.Reader.ReadString('\n')
	id = strings.TrimSpace(id)
	err := h.Service.DeleteStudent(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("刪除成功")
}
