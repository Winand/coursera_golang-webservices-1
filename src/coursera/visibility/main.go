package main

import (
	"coursera/visibility/person"
	"fmt"
)

func main() {
	p := person.NewPerson(1, "rvasily", "secret")

	// fmt.Printf("main.PrintPerson: %+v\n", p.secret) // p.secret undefined

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}
