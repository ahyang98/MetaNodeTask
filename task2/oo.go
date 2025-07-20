package task2

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func NewRectangle() Shape {
	return &Rectangle{}
}

func (r *Rectangle) Area() {
	println("In Rectangle Area")
}

func (r *Rectangle) Perimeter() {
	println("In Rectangle Perimeter")
}

type Circle struct {
}

func NewCircle() Shape {
	return &Circle{}
}

func (c *Circle) Area() {
	println("In Circle Area")
}

func (c *Circle) Perimeter() {
	println("In Circle Perimeter")
}

func TestOO() {
	var shape Shape
	shape = NewRectangle()
	shape.Area()
	shape.Perimeter()
	shape = NewCircle()
	shape.Area()
	shape.Perimeter()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Printf("EmployeeID:%d \n"+
		"Name:%s\n"+
		"Age:%d", e.EmployeeID, e.Name, e.Age)
}

func TestEmployee() {
	employee := &Employee{
		EmployeeID: 1,
		Person:     Person{"zhangsan", 100},
	}
	employee.PrintInfo()
}
