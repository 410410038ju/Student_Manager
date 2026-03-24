package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"student-manager/student" // 資料夾路徑
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== 學生管理系統 =====")
		fmt.Println("1. 新增學生")
		fmt.Println("2. 查詢學生")
		fmt.Println("3. 查詢所有學生")
		fmt.Println("4. 修改學生")
		fmt.Println("5. 刪除學生")
		fmt.Println("6. 離開")
		fmt.Print("請選擇功能: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			student.AddStudent()
		case "2":
			student.SearchStudent()
		case "3":
			student.SearchAllStudents()
		case "4":
			student.UpdateStudent()
		case "5":
			student.DeleteStudent()
		case "6":
			fmt.Println("系統結束")
			return
		default:
			fmt.Println("無效選項，請重新輸入")
		}
	}
}
