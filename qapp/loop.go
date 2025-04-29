package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

// Render is a utility to start a rendering Gio app.
func Render(fn func(ops *op.Ops)) {
	go func() {
		w := new(app.Window)
		var ops op.Ops

		for {
			e := w.Event()
			switch e := e.(type) {
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				fn(gtx.Ops)
				e.Frame(gtx.Ops)

			case app.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}

// Input is a utility to start a rendering and input Gio app.
func Input(fn func(ops *op.Ops)) {
	Render(fn)
}

// InputSize is a utility to start a rendering and input Gio app with window size.
func InputSize(fn func(ops *op.Ops, windowSize image.Point)) {
	go func() {
		w := new(app.Window)
		var ops op.Ops

		for {
			e := w.Event()
			switch e := e.(type) {
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				fn(gtx.Ops, gtx.Constraints.Max)
				e.Frame(gtx.Ops)

			case app.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}

// Metric is a utility to start a rendering, input and metrics Gio app.
func Metric(fn func(ops *op.Ops, metric unit.Metric)) {
	go func() {
		w := new(app.Window)
		var ops op.Ops

		for {
			e := w.Event()
			switch e := e.(type) {
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				fn(gtx.Ops, gtx.Metric)
				e.Frame(gtx.Ops)

			case app.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}

// Layout is a utility to start a layouting Gio app.
func Layout(lay func(gtx layout.Context) layout.Dimensions) {
	go func() {
		w := new(app.Window)
		var ops op.Ops

		for {
			e := w.Event()
			switch e := e.(type) {
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				lay(gtx)
				e.Frame(gtx.Ops)

			case app.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
