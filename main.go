package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

type Panel struct {
	name           string
	body           string
	x0, y0, x1, y1 int
}

func (w *Panel) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x0, w.y0, w.x1, w.y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, w.body)
	}
	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()
	left := &Panel{"left", "left", 0, 0, maxX/2, maxY -1}
	right := &Panel{"right", "right", maxX/2, 0, maxX - 1, maxY -1}
	g.SetManager(left, right)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}