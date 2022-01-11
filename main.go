package main

import (
	g "github.com/AllenDang/giu"
	"image/color"
	"os"
)

func main() {
	initSetting()
	app := g.NewMasterWindow("skin.ini Generator", 490, 768, g.MasterWindowFlagsNotResizable)
	app.SetBgColor(color.RGBA{R: 42, G: 21, B: 73, A: 255})
	app.Run(loop)
}

func loop() {
	g.SingleWindowWithMenuBar().Layout(
		callLayout()...,
	)

	if aboutToggle {
		g.Msgbox("About", "skin.ini Generator v0.1\nis written in Go and created by Rz.").
			Buttons(g.MsgboxButtonsOk).ResultCallback(func(r g.DialogResult) {
			aboutToggle = false
		})
	}
}

func close() {
	if edited {
		g.Msgbox("Quit without saving", "skin.ini has been modified.\nWould you like to close the application?").
			Buttons(g.MsgboxButtonsYesNo).ResultCallback(func(r g.DialogResult) {
			switch r {
			case g.DialogResultYes:
				os.Exit(0)
			case g.DialogResultNo:
				return
			}
		})
	} else {
		os.Exit(0)
	}
}
