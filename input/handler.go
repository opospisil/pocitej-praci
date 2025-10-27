// Package input handles keyboard input for the game.
package input

import (
	"github.com/opospisil/pocitej-praci/game"
)

// Handler processes keyboard input and triggers game responses.
type Handler struct {
	game *game.Game
}

// NewHandler creates a new input Handler.
func NewHandler(g *game.Game) *Handler {
	return &Handler{game: g}
}

// ProcessKeyPress handles a key press event with the given rune.
func (h *Handler) ProcessKeyPress(ch rune) {
	if game.IsCharacterKey(ch) {
		h.game.DisplayCharacter(ch)
	}
}
