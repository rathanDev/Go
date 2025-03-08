package two

import "fmt"

type Student struct {
	Name string
	Gender string
}

func mainO() {

	var stPointer *Student
	stPointer = new(Student)
	stPointer.Name = "n1"
	fmt.Println(stPointer)

	var st Student
	st.Name = "n2"
	st.Gender = "M"
	fmt.Println(st)

	st3 := Student{}
	st3.Name = "n3"
	fmt.Println(st3)

	st4 := Student{"n4", "M"}
	fmt.Println(st4)
}

func updateStudent(st *Student) {

}