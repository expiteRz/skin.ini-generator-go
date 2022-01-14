package main

import (
	g "github.com/AllenDang/giu"
	"github.com/faiface/mainthread"
	"image/color"
	"os"
)

func main() {
	initSetting()
	app := g.NewMasterWindow("skin.ini Generator", 490, 768, g.MasterWindowFlagsNotResizable)
	app.SetBgColor(color.RGBA{R: 42, G: 21, B: 73, A: 255})
	app.Run(loop)
	mainthread.Run(fn)
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

	if errorBox {
		g.Msgbox("Error", errorMsg).ResultCallback(func(r g.DialogResult) {
			errorBox = false
		})
	}

	if g.IsKeyDown(g.KeyLeftControl+g.KeyO) || g.IsKeyDown(g.KeyRightControl+g.KeyO) {
		openFile()
	} else if g.IsKeyDown(g.KeyLeftControl+g.KeyN) || g.IsKeyDown(g.KeyRightControl+g.KeyN) {
		initSetting()
	} else if g.IsKeyDown(g.KeyLeftControl+g.KeyS) || g.IsKeyDown(g.KeyRightControl+g.KeyS) {
		// TODO Saving file process
	}
}

func fn() {
}

func closeFunc() {
	if edited {
		g.Msgbox("Quit without saving", "skin.ini has been modified.\nWould you like to closeFunc the application?").
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
