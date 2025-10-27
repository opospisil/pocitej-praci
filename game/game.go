package game

import (
	"time"

	"fyne.io/fyne/v2"
	"github.com/opospisil/pocitej-praci/config"
)

// Game manages the character display and timing.
type Game struct {
	renderer  *Renderer
	container fyne.CanvasObject
	hideTimer *time.Timer
}

// NewGame creates a new Game with the given configuration.
func NewGame(cfg *config.Config, displayContainer fyne.CanvasObject) *Game {
	return &Game{
		renderer:  NewRenderer(cfg),
		container: displayContainer,
	}
}

// DisplayCharacter shows a character for the configured duration.
func (g *Game) DisplayCharacter(ch rune) {
	// Cancel any pending hide timer
	if g.hideTimer != nil {
		g.hideTimer.Stop()
	}

	// Create and display the character
	display := g.renderer.CreateCharacterDisplay(ch)
	if cont, ok := g.container.(*fyne.Container); ok {
		cont.RemoveAll()
		cont.Add(display)
		cont.Refresh()
	}

	// Schedule hiding after the configured duration
	g.hideTimer = time.AfterFunc(time.Duration(g.renderer.cfg.DisplayDurationMs)*time.Millisecond, func() {
		if cont, ok := g.container.(*fyne.Container); ok {
			cont.RemoveAll()
			fyne.Do(func() {
				cont.Refresh()
			})
		}
	})
}
