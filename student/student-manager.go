package student

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Student 結構
type Student struct {
	ID    string
	Name  string
	Age   int
	Class string
}

// 全域 slice 存學生資料
var Students []Student

// 新增學生
func AddStudent() {
	reader := bufio.NewReader(os.Stdin)
	var s Student

	fmt.Print("輸入學生ID: ")
	s.ID, _ = reader.ReadString('\n')
	s.ID = strings.TrimSpace(s.ID)

	fmt.Print("輸入學生姓名: ")
	s.Name, _ = reader.ReadString('\n')
	s.Name = strings.TrimSpace(s.Name)

	fmt.Print("輸入學生年齡: ")
	fmt.Scanf("%d\n", &s.Age)

	fmt.Print("輸入學生班級: ")
	s.Class, _ = reader.ReadString('\n')
	s.Class = strings.TrimSpace(s.Class)

	Students = append(Students, s)
	fmt.Println("學生新增成功！")
}

// 查詢學生
func SearchStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("請輸入學生ID查詢: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for _, s := range Students {
		if s.ID == id {
			fmt.Printf("學生ID: %s\n姓名: %s\n年齡: %d\n班級: %s\n", s.ID, s.Name, s.Age, s.Class)
			return
		}
	}
	fmt.Println("找不到學生")
}

// 查詢所有學生
func SearchAllStudents() {
	if len(Students) == 0 {
		fmt.Println("目前沒有學生資料")
		return
	}
	for _, s := range Students {
		fmt.Printf("學生ID: %s\n姓名: %s\n年齡: %d\n班級: %s\n", s.ID, s.Name, s.Age, s.Class)
	}
}

// 修改學生
func UpdateStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("請輸入學生ID修改: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	for i, s := range Students {
		if s.ID == id {
			fmt.Print("輸入新的學生姓名: ")
			s.Name, _ = reader.ReadString('\n')
			s.Name = strings.TrimSpace(s.Name)

			fmt.Print("輸入新的學生年齡: ")
			fmt.Scanf("%d\n", &s.Age)

			fmt.Print("輸入新的學生班級: ")
			s.Class, _ = reader.ReadString('\n')
			s.Class = strings.TrimSpace(s.Class)

			Students[i] = s
			fmt.Println("學生資料修改成功！")
			return
		}
	}
	fmt.Println("找不到學生")
}

// 刪除學生
func DeleteStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("請輸入學生ID刪除: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	for i, s := range Students {
		if s.ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			fmt.Println("學生資料刪除成功！")
			return
		}
	}
	fmt.Println("找不到學生")
}
