package main

import (
	"encoding/json"
	"time"
)

func main() {
	// don't care about the concrete type, only print it
	myHomework := &Homework{
		Subject:  "Math",
		Deadline: time.Now(),
		Content:  "Give a star to the repository",
	}

	yourHomework := myHomework.Clone()

	mb, _ := json.Marshal(myHomework)
	yb, _ := json.Marshal(yourHomework)

	println(string(mb), string(yb))

	// type assertion and then get the concrete type
	dream := &Dream{Content: "I want to be a helpful person"}
	dream2 := dream.Clone()

	d, ok := dream2.(*Dream)
	if ok {
		println(d.Content)
	} else {
		println(dream2)
	}

}
