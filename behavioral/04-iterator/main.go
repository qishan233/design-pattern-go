package main

func main() {
	a := GiveMeAggregate()

	println("Normal iterator iterate begin:")
	normalIterator := a.CreateIterator(NormalIteratorType)
	for normalIterator.HasNext() {
		item := normalIterator.Next()
		if item != nil {
			println(item.Name)
		}
	}

	println("Odd iterator iterate begin:")
	oddIterator := a.CreateIterator(OddIteratorType)
	for oddIterator.HasNext() {
		item := oddIterator.Next()
		if item != nil {
			println(item.Name)
		}
	}
}
