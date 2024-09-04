package main

type WidgetFactory interface {
	CreateScrollBar() ScrollBar
	CreateWindow() Window
}

func GetWidgetFactory(factoryType string) WidgetFactory {
	if factoryType == "PM" {
		return &PMWidgetFactory{}
	} else if factoryType == "Motif" {
		return &MotifWidgetFactory{}
	}
	return nil
}

type PMWidgetFactory struct {
}

func (p *PMWidgetFactory) CreateScrollBar() ScrollBar {
	return &PMScrollBar{}
}

func (p *PMWidgetFactory) CreateWindow() Window {
	return &PMWindow{}
}

type MotifWidgetFactory struct{}

func (m *MotifWidgetFactory) CreateScrollBar() ScrollBar {
	return &MotifScrollBar{}
}

func (m *MotifWidgetFactory) CreateWindow() Window {
	return &MotifWindow{}
}
