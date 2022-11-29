package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p1 := person{"Alex", "Anderson", contactInfo{"alex.anderson@example.com", 1000}}

	p2 := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email: "john.do@example.com",
			zip:   2000,
		},
	}

	var p3 person
	p3.firstName = "Will"
	p3.lastName = "Smith"
	p3.contactInfo.email = "will.smith@example.com"
	p3.contactInfo.zip = 3000

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	p1.updateName("Jimmy")
	fmt.Println(p1)
}

/* implement Stringer interface for contectInfo */
func (ci contactInfo) String() string {
	return fmt.Sprintf("%v %v", ci.email, ci.zip)
}

/* implement Stringer interface for person */
func (p person) String() string {
	return fmt.Sprintf("%v %v: %v", p.firstName, p.lastName, p.contactInfo)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
