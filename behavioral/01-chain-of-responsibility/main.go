package main

func main() {
	handler := GetHandler()

	// client does not know which handler will handle the request. that means we can change the dispatch logic at runtime
	handler.HandleRequest("A")
	handler.HandleRequest("B")
	handler.HandleRequest("C")
}
