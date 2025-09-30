package main

import "fmt"

type Thing struct {
	ThingInt int
}

func (t Thing) DoThing() {}

func DoThingWithTHing(t Thing) {}

func main() {
	recieveOnly := make(chan<- int, 0)
	defer func() {
		if e := recover(); e != nil {
			// handle panic in e
		}
	}()
	fmt.Print
	fmt.Println

	t := new(Thing) // nil of type Thing
	t := &Thing{}   // instance of Thing
	panic("wat")

	t.DoThing()
	t.ThingInt
	(*t).DoThing()
	(*t).ThingInt

	DoThingWithTHing(*t)

}
