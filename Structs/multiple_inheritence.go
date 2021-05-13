package main
import "fmt"
type Student struct {
	FName string
	LName string
}
type Exams struct {
	UnitTestsExam int
}
func(exam Exams) TakeExam(){
fmt.Println("exam")
}
func(std Student) TakeExam(){
	fmt.Println("student")
}
type SchoolStd struct {
	Student
	Exams
}
func main() {
	std := Student{"Sachin", "T"}
	exam := Exams{10}
	schoolstd := SchoolStd{std, exam}
	fmt.Println(schoolstd.FName, schoolstd.UnitTestsExam)
	schoolstd.Student.TakeExam()
	schoolstd.Exams.TakeExam()
}