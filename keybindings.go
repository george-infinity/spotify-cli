package main

import (
	"github.com/jroimartin/gocui"
)

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, moveCursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, moveCursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, runCommand); err != nil {
		return err
	}

	return nil
}

func moveCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		return MoveCursor(g, v, -1)
	}

	return nil
}

func moveCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		return MoveCursor(g, v, 1)
	}

	return nil
}

func runCommand(g *gocui.Gui, v *gocui.View) error {
	return RunCommand(g, v)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
