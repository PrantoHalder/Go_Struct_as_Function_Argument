package main

import "fmt"

type Person struct {
	Name string
	Age int 
	subject string
	marks int
}

func struckAsFunctionArgument () {
fmt.Println("this code is to use struch as function argument")
var person Person
var person1 Person
person.Name = "Pranto"
person.Age = 25
person.subject = "CSE"
person.marks = 90

person1.Name = "shovon"
person1.Age = 26
person1.subject = "EEE"
person1.marks = 82

printFunction(person,person1)
}

func printFunction (p , q Person){
  fmt.Println(p)
  fmt.Println(q)
}