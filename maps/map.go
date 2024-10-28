package main

import "fmt"

func addStudent(grades map[string]int, name string, grade int) {
	grades[name] = grade
}
func main() {
	ages := map[string]int{
		"ram":      18,
		"prashant": 20,
		"kevin":    17,
	}
	fmt.Println("the lenght of the ages array is", len(ages))
	fmt.Println("an ages array", ages)

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	//with func addStudent
    studentGrades := make(map[string]int)
    addStudent(studentGrades, "Alice", 90)
    addStudent(studentGrades, "Bob", 85)

    fmt.Println("Student Grades:", studentGrades)
}
