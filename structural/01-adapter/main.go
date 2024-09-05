package main

func main() {
	classAdapter := NewClassAdapter()
	Request("class adapter:", classAdapter)

	objectAdapter := NewObjectAdapter(&AdapteeImpl{})
	Request("object adapter:", objectAdapter)
}

func Request(tips string, t Target) {
	println(tips, t.Request())
}
