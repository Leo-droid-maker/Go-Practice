package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName, lastName string
	// contact ContactInfo
	ContactInfo
}

func main() {
	// alex := Person{"Alex", "Anderson"} - BAD choise to declare struct object
	// alex := Person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName) - better choise to declare an object
	var alex Person
	fmt.Println(alex)
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v\n", alex)

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v\n", jim.ContactInfo)
	jim.print()
	// jimPointer := &jim // - Link to memory adress of the value this variable

	// jimPointer.updateFirstName("jimmy")
	jim.updateFirstName("jimmy")
	jim.print()

	mySlice := []string{"How", "Are", "You?"}

	updateSlice(mySlice)
	/* <- Here's another copy of the slice, but the pointer is on the same
	adress with same value !!!
	And 'cause slices
	can grow and shrink
	and used 99% of the time for lists of elements*/

	fmt.Println(mySlice) // [Bye Are You?]

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (pointerToPerson *Person) updateFirstName(newFirstName string) {
	/* Give me value this memory adress
	 *Person - just description
	 *pointerToPerson - operator */
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() { // struct method with resiver
	fmt.Printf("%+v\n", p)
}
