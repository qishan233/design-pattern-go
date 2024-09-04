package main

func main() {
	thing1 := GetTheMostImportantThing()
	println(thing1.GiveMeGuide())

	thing2 := GetTheMostImportantThing()
	println(thing2.GiveMeGuide())

	println("there is only one most import thing:", thing1 == thing2)
}
