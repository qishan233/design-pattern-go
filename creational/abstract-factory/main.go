package main

func main() {
	systemType := "PM"

	widgetFactory := GetWidgetFactory(systemType)

	// window and scroll bar are created by the same factory to make sure they are compatible
	// client code does not need to know the concrete classes but the abstract factory
	window := widgetFactory.CreateWindow()
	window.Draw()

	scrollBar := widgetFactory.CreateScrollBar()
	scrollBar.Scroll()
}
