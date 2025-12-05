package main

import "fmt"

func main() {
	//Ein Schüler (Student) hat einen Vor- und Nachnamen.
	//Eine Klasse (Class) besteht aus einer Reihe von Schülern.
	//Ein Modul hat eine eindeutige Nummer (z.B. 346) und wird von einer Reihe von Klassen besucht.
	//Erstelle die notwendigen Datenstrukturen mit entsprechenden Beispieldaten (d.h. mindestens zwei Klassen mit je drei Schülern und mindestens drei Module, die von einer oder mehreren Klassen besucht werden). Gib die Daten anschliessend per fmt.Println() auf die Konsole aus.
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	var studentOne = Student{
		"Max",
		"Mustermann",
	}
	studentTwo := Student{
		"Clarissa",
		"Fairchild",
	}
	studentThree := Student{
		"Clint",
		"Barton",
	}
	studentFour := Student{
		"Lilo",
		"Pelekai",
	}
	studentFive := Student{
		"Lucy",
		"Chen",
	}
	studentSix := Student{
		"Luka",
		"Couffaine",
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class []Student

	classOne := Class{studentOne, studentThree, studentFour}
	classTwo := Class{studentTwo, studentFive, studentSix}
	// TODO: declare a map of modules being attended by multiple classes
	var modules = map[uint][]Class{
		104: {classTwo},
		117: {classOne},
		364: {classOne, classTwo},
	}
	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
