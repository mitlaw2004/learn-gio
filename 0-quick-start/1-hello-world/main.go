// SPDX-License-Identifier: Unlicense OR MIT
package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"             // app contains Window handling.
	"gioui.org/op"              // op is used for recording different operations.
	"gioui.org/text"            // text contains constants for text layouting.
	"gioui.org/widget/material" // material contains material design widgets.
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	// theme contains constants for theming.
	theme := material.NewTheme()
	// ops will be used to encode different operations
	var ops op.Ops
	for {
		// detect the type of the event.
		switch e := window.Event().(type) {
		// this is sent when the application should re-render.
		case app.DestroyEvent:
			return e.Err
			// this is sent when the application should re-render.
		case app.FrameEvent:
			// This graphics context is used for managing the rendering and event information.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
