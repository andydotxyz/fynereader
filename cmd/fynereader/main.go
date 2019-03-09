// Package main launches the imgur picture browser example directly
package main

import (
	"fyne.io/fyne/app"
	"github.com/andydotxyz/fynereader"
)

func main() {
	app := app.New()

	fynereader.Show(app)
	app.Run()
}
