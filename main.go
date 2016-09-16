package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"log/syslog"
)

var (
	logWriter *syslog.Writer
)

func main() {
	logWriter, err := syslog.New(syslog.LOG_NOTICE, "spotify-cli")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logWriter)

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	if err := keybindings(g); err != nil {
		log.Fatal(err)
	}

	g.SetLayout(layout)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("controls", 0, 0, 20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		err = g.SetCurrentView("controls")
		if err != nil {
			return err
		}

		v.Highlight = true
		v.Wrap = true
		v.Editable = false
		v.Autoscroll = false

		for _, item := range menu {
			fmt.Fprintln(v, item.text)
		}
	}

	if v, err := g.SetView("status", 20, 0, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Editable = false
		v.Autoscroll = false

		status, err := GetStatus()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(v, "Artist: ", status.artist)
		fmt.Fprintln(v, "Album: ", status.album)
		fmt.Fprintln(v, "Track: ", status.track)
	}

	return nil
}

func MoveCursor(g *gocui.Gui, v *gocui.View, my int) error {
	cx, cy := v.Cursor()
	ny := cy + my

	if ny < 0 || ny >= len(menu) {
		return nil
	}

	if err := v.SetCursor(cx, ny); err != nil {
		return err
	}

	return nil
}

func RunCommand(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	return menu[cy].fn(g, v)
}
