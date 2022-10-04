package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name       string
	Address    string
	Job        string
	Motivation string
}

func main() {

	var pers1 Person
	var pers2 Person

	pers1.Name = "Rizky"
	pers1.Address = "Jakarta"
	pers1.Job = "Programmer"
	pers1.Motivation = "I want to be a programmer"

	pers2.Name = "Anis"
	pers2.Address = "Bandung"
	pers2.Job = "backend"
	pers2.Motivation = "I want to be a backend developer"

	if os.Args[1] == "1" {
		printPerson(pers1)
	} else if os.Args[1] == "2" {
		printPerson(pers2)
	} else {
		fmt.Println("Input is not valid")
	}

}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.Name)
	fmt.Println("Address: ", pers.Address)
	fmt.Println("Job: ", pers.Job)
	fmt.Println("Motivation: ", pers.Motivation)
}
