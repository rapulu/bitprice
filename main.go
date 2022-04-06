package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


func main() {
	fmt.Println("Starting app...")


	a := app.New()

	w := a.NewWindow("BitPrice App")
	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}