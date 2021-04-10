package university

import (
	"fmt"
	"time"

	"github.com/arthurh0812/university/country"
)

type Person struct {
	firstName string
	lastName string
	dateOfBirth time.Time
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) DateOfBirth() time.Time {
	return p.dateOfBirth
}

func NewPerson(firstName, lastName string, dob time.Time) *Person {
	p := &Person{
		firstName: firstName,
		lastName: lastName,
		dateOfBirth: dob,
	}

}

type Address struct {
	country country.Country
	city string
	streetName string
	postalCode int32
}

func (a Address) Country() country.Country {
	return a.country
}

func (a Address) City() string {
	return a.city
}

func (a Address) StreetName() string {
	return a.streetName
}

func (a Address) PostalCode() int32 {
	return a.postalCode
}