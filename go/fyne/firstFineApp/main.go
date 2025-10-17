package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("First Fyne App")

	clockBinding := binding.NewString()
	clock := widget.NewLabelWithData(clockBinding)

	rect := canvas.NewRectangle(color.Black)
	window.SetContent(rect)

	window.Resize(fyne.NewSize(150, 100))

	go func() {
		for t := range time.Tick(time.Second) {
			clockBinding.Set("Current time: " + t.Format("15:04:05"))
		}
	}()

	window.SetContent(clock)
	window.ShowAndRun()
}
