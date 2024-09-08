package main

func main() {
	println("Story begins with a requirement document")
	pm := GiveMeProductManager()

	pm.WriteRequirementDocument()
}