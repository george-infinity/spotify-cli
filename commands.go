package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"os/exec"
)

type Command struct {
	text string
	fn   func(*gocui.Gui, *gocui.View) error
}

type Status struct {
	artist string
	album  string
	track  string
}

var menu = []Command{
	Command{"Play/Pause", CmdPlayPause},
	Command{"Next Track", CmdNextTrack},
	Command{"Previous Track", CmdPrevTrack},
}

func CmdPlayPause(g *gocui.Gui, v *gocui.View) error {
	log.Print("Play/Pause")

	return nil
}

func CmdNextTrack(g *gocui.Gui, v *gocui.View) error {
	log.Print("Next Track")

	return nil
}

func CmdPrevTrack(g *gocui.Gui, v *gocui.View) error {
	log.Print("Previous Track")

	return nil
}

func GetStatus() (Status, error) {
	s := Status{}

	artist, err := RunOsascript("artist of current track as string")
	if err != nil {
		return s, err
	}

	s.artist = artist

	return s, nil
}

func RunOsascript(args string) (string, error) {
	log.Print("osascript -e 'tell application \"Spotify\" to " + args + "'")

	cmd := exec.Command("osascript", "-e", "'tell application \"Spotify\" to "+args+"'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
