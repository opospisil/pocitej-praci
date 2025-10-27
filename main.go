package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/opospisil/pocitej-praci/config"
	"github.com/opospisil/pocitej-praci/game"
	"github.com/opospisil/pocitej-praci/input"
)

func main() {
	cfg := config.NewConfig()
	myApp := app.New()
	myWindow := myApp.NewWindow("Character Display Game")

	// Create a black background rectangle
	background := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 255})

	// Create a center container for displaying characters
	displayContainer := container.NewCenter()

	// Create a container with background and overlay the display container
	mainContainer := container.NewStack(background, displayContainer)

	// Initialize game
	g := game.NewGame(cfg, displayContainer)

	// Initialize input handler
	handler := input.NewHandler(g)

	// Set the main container as the window content
	myWindow.SetContent(mainContainer)

	// Set up keyboard input handling for typed runes on the canvas
	myWindow.Canvas().SetOnTypedRune(func(ch rune) {
		handler.ProcessKeyPress(ch)
	})

	// Make the window large (approximate fullscreen)
	myWindow.Resize(fyne.NewSize(1920, 1080))

	myWindow.ShowAndRun()
}
