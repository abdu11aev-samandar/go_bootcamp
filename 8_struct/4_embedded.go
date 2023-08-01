package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John",
		salary: 4000,
		contactInfo: Contact{
			email:   "1233@da.com",
			address: "Street 20, London",
			phone:   03030300303,
		},
	}
	fmt.Printf("%+v\n", john)

	fmt.Printf("Employee's email:%s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@the.com"
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email)

	myContact := Contact{email: "anri@dom.com", phone: 1223231, address: "Andijon"}
	fmt.Println(myContact)
}
