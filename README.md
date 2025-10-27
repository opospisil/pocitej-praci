# Pocitej Praci (Compute work) - Interactive Game for Toddlers

A fullscreen interactive application designed for toddlers that creates engaging visual feedback when they press character keys on the keyboard.

## Motivation
My 2.5 y.o. son sees me sitting every day at the computer and working. So he started to do the same, always saying that he needs to compute work (pocitat praci in czech)
So there he has it, now he can interact with a computer on a level that is just above what he knows. Easily asociating some letters with words while not being able to fully
understand others.


## Features

- **Character Display**: Any character key press (letters, numbers, symbols) displays the character large and centered on the screen
- **Visual Learning**: Shows both uppercase and lowercase letters for alphabetic keys
- **Color Feedback**: Each character appears in a randomly selected color
- **Auto-Hide**: Characters automatically disappear after a configurable duration (default: 1.5 second)
- **Smart Input**: Ignores non-character keys (arrows, function keys, etc.) to keep interactions simple

![Project Demo](demo.gif)


## Quick Start

### Prerequisites
- Go 1.24 or higher
- Build on Void linux with X11 development libraries:
```bash
libX11-devel \
libXrandr-devel \
libXinerama-devel \
libXcursor-devel \
libXext-devel \
libxkbcommon-devel \
libXxf86vm-devel
```
### Build
```bash
go build -o pocitej-praci .
```

### Run
```bash
./pocitej-praci
```
## Configuration

Edit `config/config.go` to customize:

| Setting | Default | Purpose |
|---------|---------|---------|
| `DisplayDurationMs` | 1500 | How long to show each character (milliseconds) |
| `FontSize` | 300 | Text size in points  |
| `ColorPalette` | 6 colors | Vibrant colors for visual feedback |

[License](LICENSE)
