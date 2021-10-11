// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Example - Demonstrates debug camera. Use "t" to toggle camera modes.
package main

import (
	"github.com/waryway/go2dgame/internal/controller"
	//"github.com/waryway/go2dgame/internal/models"
	//"image"
	"log"
	//"math"

	"azul3d.org/engine/gfx"
	//"azul3d.org/engine/gfx/camera"
	//"azul3d.org/engine/gfx/gfxutil"
	"azul3d.org/engine/gfx/window"
	//"azul3d.org/engine/lmath"
	//"azul3d.org/examples/abs"
)

func main() {
	// Routine for the Main Coordinates
	go func() {
		props := window.NewProps()
		props.SetTitle("Main camera {FPS}")
		props.SetPos(0, 100)
		props.SetSize(640, 400)
		w, r, err := window.New(props)
		if err != nil {
			log.Fatal(err)
		}

		gfxLoopWindow1(w, r)
	}()

	// Enter the main loop.
	window.MainLoop()
}

// gfxLoopWindow1 runs the main camera window
func gfxLoopWindow1(w window.Window, d gfx.Device) {
	var objects []controller.Object
	cubes := controller.Object{}
	cubes.Init("cube")
	objects = append(objects, cubes)

	var tripod controller.Tripod
	tripod.Init(d, objects)
	controller.KeyboardListener(w, d, tripod)
}
