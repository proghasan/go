package main

import "fmt"

type Artist struct {
	Name  string
	Genre string
	Songs int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {

	me := &Artist{
		Name:  "James",
		Genre: "BB Band",
		Songs: 10,
	}

	fmt.Printf("%s released their %d th song \n ", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
