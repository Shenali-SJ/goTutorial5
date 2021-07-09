package main

import "fmt"

type student struct {
	firstName string
	lastName string
	age int
}

func filterStudent(s []student, f func(student) bool) []student {
	var students []student

	for _, v := range s {
		if f(v) == true {
			students = append(students, v)
		}
	}
	return students
}

func main() {
	s1 := student{
		firstName: "rachel",
		lastName:  "green",
		age:       19,
	}
	s2 := student{
		firstName: "monica",
		lastName:  "geller",
		age:       19,
	}
	s3 := student{
		firstName: "phoebe",
		lastName:  "buffay",
		age:       22,
	}

	students := []student {s1, s2, s3}

	f := filterStudent(students, func(s student) bool {
		if s.age < 20 {
			return true
		}
		return false
	})

	fmt.Println(f)
}
