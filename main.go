package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func main() {
	os.Setenv("FYNE_DISABLE_GL", "true")
	os.Setenv("FYNE_SCALE", "1")

	myApp := app.New()
	myWindow := myApp.NewWindow("Parameter Manager")

	label := widget.NewLabel("Hello, 파라미터 관리 프로그램!")

	myWindow.SetContent(container.NewVBox(
		label,
	))

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
