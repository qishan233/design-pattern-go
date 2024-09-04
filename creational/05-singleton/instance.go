package main

import "sync"


type TheMostImportantThing struct {
	content string
}

func (t *TheMostImportantThing) GiveMeGuide() string {
	return t.content
}

var (
	once     sync.Once
	instance *TheMostImportantThing
)

// GetTheMostImportantThing the global entry point to get the most important thing
func GetTheMostImportantThing() *TheMostImportantThing {
	once.Do(func() {
		instance = &TheMostImportantThing{content: "Live a good life and do things that matter"}
	})

	return instance
}
