package main

type ObjectAdapter struct {
	Adaptee Adaptee
}

func NewObjectAdapter(adaptee Adaptee) *ObjectAdapter {
	return &ObjectAdapter{
		Adaptee: adaptee,
	}
}

func (a *ObjectAdapter) Request() string {
	if a.Adaptee == nil {
		return "cant't finish request without Adaptee"
	}

	return a.Adaptee.SpecificRequest() + ", object adapter do nothing"
}

type AdapteeImpl struct{}

func (a *AdapteeImpl) SpecificRequest() string {
	return "AdapteeImpl SpecificRequest"
}
