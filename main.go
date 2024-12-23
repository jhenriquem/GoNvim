package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	gb "github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/modes"
	"github.com/jhenriquem/go-neovim/render"
	"github.com/jhenriquem/go-neovim/screen"
)

func main() {
	log.Println("Iniciando o programa...")

	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	_, ScreenHeight := screen.Screen.Size()

	if gb.CurrentLine < gb.ScrollOffSet {
		gb.ScrollOffSet = gb.CurrentLine
	}

	if gb.CurrentLine >= gb.ScrollOffSet+ScreenHeight-1 {
		gb.ScrollOffSet = gb.CurrentLine - (ScreenHeight - 1)
	}

	// Inicialização de lines
	gb.Lines = make([][]rune, 1)

	for gb.RunningApp {
		render.RenderLines(ScreenHeight)
		render.RenderStatusLine()
		render.RenderCommandLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:
			if modes.CurrentMODE == "NORMAL" {
				modes.KeymapsEventsForNormalMode(ev)
			} else if modes.CurrentMODE == "INSERT" {
				modes.KeymapsEventsForInsertMode(ev)
			}
		}
	}
}
