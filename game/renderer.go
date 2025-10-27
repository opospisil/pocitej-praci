package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/opospisil/pocitej-praci/config"
)

// Renderer handles the visual display of characters.
type Renderer struct {
	cfg *config.Config
}

// NewRenderer creates a new Renderer with the given configuration.
func NewRenderer(cfg *config.Config) *Renderer {
	return &Renderer{cfg: cfg}
}

// GetRandomColor returns a random color from the palette.
func (r *Renderer) GetRandomColor() color.Color {
	if len(r.cfg.ColorPalette) == 0 {
		return color.White
	}
	return r.cfg.ColorPalette[rand.Intn(len(r.cfg.ColorPalette))]
}

// CreateCharacterDisplay creates a Fyne container displaying the character(s).
func (r *Renderer) CreateCharacterDisplay(ch rune) fyne.CanvasObject {
	col := r.GetRandomColor()
	text := string(ch)

	// For letters, show both upper and lower case
	if unicode.IsLetter(ch) {
		text = fmt.Sprintf("%c %c", unicode.ToUpper(ch), unicode.ToLower(ch))
	}

	richText := canvas.NewText(text, col)
	richText.TextSize = float32(r.cfg.FontSize)
	richText.TextStyle.Bold = true
	richText.Alignment = fyne.TextAlignCenter

	// Create a container to center the text
	return container.NewCenter(richText)
}

// IsCharacterKey returns true if the key is a writable character.
func IsCharacterKey(ch rune) bool {
	// Accept letters, digits, and common punctuation/symbols
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) ||
		strings.ContainsRune("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", ch)
}
