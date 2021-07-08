package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	type Team struct {
		name    string
		players [3]Person
	}
	var myTeam Team
	fmt.Println((myTeam))

	players := [...]Person{Person{name: "Forrest", age: 45}, Person{name: "Gordon", age: 34}, Person{name: "Terry", age: 56}}
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)

}
