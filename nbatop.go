package main

import (
	"log"
	
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"schedule"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello, world."
	p.SetRect(0, 0, 25, 5)

	l := schedule()
	ui.Render(l)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
