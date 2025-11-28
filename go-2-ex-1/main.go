package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	BirthDay   byte
	BirthMonth byte
	BirthYear  uint16
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	BornOn           BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			"Sarah",
			"Christen",
		},
		BornOn: BirthDate{
			30,
			8,
			2009,
		},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u264D', // TODO: adjust
	}
	fmt.Println(me)                   //printing struct like this doesn't work for zodiac sign emoji
	fmt.Printf("%c\n", me.ZodiacSign) //but it works like that

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
