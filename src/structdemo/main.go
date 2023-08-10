package main
import (
	"fmt"
	"encoding/json"
)

type Employee struct {
	ID int
	Firstname string
	Lastname string
	Address string
}

type Person struct {
	ID int 
	Firstname string `json:"name"` // 输出json时将Firstname改为name
	Lastname string
	Address string `json:"address,omitempty` // 输出json时，如果address为空，则不输出
}

type Student struct {
	Infomation Person
	StudentId int 
}

type Teacher struct {
	Person
	TeacherId int
}

var test Employee
func main(){
	employee := Employee{1,"dudy","zhao","杭州"}
	fmt.Println("employee",employee)
	
	employee1:= Employee{Lastname:"zhang",Address:"china"}
	fmt.Println("employee1",employee1)
	employeeCopy := &employee
	employeeCopy.Lastname = "wang"
	fmt.Println("employeeCopy",employee)

	var student Student
	student.Infomation.Firstname = "dudy"
	fmt.Println("student",student)

	teacher := Teacher{
		Person: Person{
			Firstname: "judy",
		},
	}
	teacher.Lastname = "wang"
	fmt.Println("teacher",teacher)

	data,_ := json.Marshal(teacher)
	fmt.Printf("%s\n",data)

	var teacheStruct Teacher
	json.Unmarshal(data,&teacheStruct)
	fmt.Printf("%v",teacheStruct)
}