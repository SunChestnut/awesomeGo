package student

import (
	"fmt"
)

type student struct {
	name          string
	age           int
	school        string
	classAndGrade string
}

func createStudent(name, school string) student {
	return student{name: name, school: school}
}

func main() {

	var personA student
	personA = student{name: "炮姐"}

	//personB := new(student)
	var personB student
	personB.age = 188
	personB.school = "天山常随研究所"

	personC := student{age: 456, classAndGrade: "四年B班"}

	personD := student{"D 同学", 16, "家里蹲超长寿大学", "三年A班"}

	students := []student{
		personA,
		personB,
		personC,
		personD,
	}

	fmt.Println(students)

	personABC := createStudent("ABC", "567")
	fmt.Println(personABC)

	// classAndGrade 的默认值为空串，在上述代码中也是打印出来的，只不过不明显
	fmt.Printf("classAndGrade = %q", personABC.classAndGrade)
}
