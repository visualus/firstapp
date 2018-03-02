package main

import "fmt"
// import "reflect"

func main() {
	var msg string
	message := "This is a message"
	msg = "This is a message as well"

	fmt.Println("Hello !")
	fmt.Println(message)
	fmt.Println(msg)

	// var n bool = true
	n := 1 == 1
	m := 1 == 2

	fmt.Println(n)
	fmt.Println(m)

	grades := [3]int{97, 85, 93}
	// gradestwo := [...]int{97, 85, 93}

	fmt.Printf("Grades: %v \n", grades)

	var students [3]string
	students[0] = "Lisa"
	students[1] = "Arnold"
	students[2] = "John"

	newstudents := []string{}
	newstudents = append(newstudents, "Mark")

	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Students: %v \n", newstudents)

	fmt.Printf("Students: %v \n", len(newstudents))
	fmt.Printf("Students: %v \n", cap(newstudents))

	// MAPS and Structs

	dictionary := map[string]int{
		"California": 1,
		"Texas":      2,
		"New York":   3,
	}

	fmt.Printf("My Dict: %v \n", dictionary)
	_, ok := dictionary["California"]
	fmt.Println(ok)

	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "John Doe",
		companions: []string{
			"Liz",
			"Jo",
			"Sarah",
		},
	}

	fmt.Printf("%v", aDoctor)

	bDoctor := struct{ name string }{name: "John Dumbledore"}

	fmt.Printf("%v", bDoctor)

	type Animal struct {
		Name   string `required max:"100"`
		Origin string
	}

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)

	if true && true {
		fmt.Println("this is trues")
	}

	for i := 0; i<5; i++{
		fmt.Println(i)
	}

	s := []int{1,2,3}

	for k,v := range s{
		fmt.Println(k,v)
	}


	st := "Hello world"

	for k,v := range st{
		fmt.Println(k,string(v))
	}


	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	// deffered - executes after the main faunction

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")



}
