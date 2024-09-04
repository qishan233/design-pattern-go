package main

func main() {
	systemType := "PM"
	// systemType := "Motif"

	// get the abstract factory. you can change the systemType to decide which factory to use, and don't need to change the scroll bar and window creation code
	widgetFactory := GetWidgetFactory(systemType)

	if widgetFactory == nil {
		panic("Unsupported factory type")
	}

	// window and scroll bar are created by the same factory to make sure they are compatible
	// client code does not need to know the concrete classes but the abstract factory
	window := widgetFactory.CreateWindow()
	window.Draw()

	scrollBar := widgetFactory.CreateScrollBar()
	scrollBar.Scroll()
}
