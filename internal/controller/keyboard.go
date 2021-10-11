package controller

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/window"
	"azul3d.org/engine/keyboard"
	"azul3d.org/engine/lmath"
	"azul3d.org/engine/mouse"
	"log"
)

func KeyboardListener(w window.Window, d gfx.Device, tripod Tripod) {
	var AxisMovement = make(chan lmath.Vec3)
	var AxisRotation = make(chan lmath.Vec3)

	movement := lmath.Vec3{0, 0, 0}
	rotation := lmath.Vec3{0, 0, 0}
	go tripod.Movement(AxisMovement)
	go tripod.Rotation(AxisRotation)

	events := make(chan window.Event, 256)
	w.Notify(events, window.AllEvents)

	for {
		select {
		case e := <-events:

			switch ev := e.(type) {
			case mouse.ButtonEvent:
				log.Println(mouse.Down.String())
				break

			case keyboard.ButtonEvent:
				go func() {
					for w.Keyboard().Down(keyboard.W) {
						movement.Z = -.1
						AxisMovement <- movement
					}
					for w.Keyboard().Down(keyboard.S) {
						movement.Z = .1
						AxisMovement <- movement
					}

					for w.Keyboard().Down(keyboard.A) {
						rotation.Y = 0.05
						AxisRotation <- rotation
					}
					for w.Keyboard().Down(keyboard.D) {
						rotation.Y = -0.05
						AxisRotation <- rotation
					}

					for w.Keyboard().Down(keyboard.Q) {
						movement.X = -.1
						AxisMovement <- movement
					}
					for w.Keyboard().Down(keyboard.E) {
						movement.X = .1
						AxisMovement <- movement
					}

					for w.Keyboard().Down(keyboard.X) {
						movement.Y = -.1
						AxisMovement <- movement
					}
					for w.Keyboard().Down(keyboard.Z) {
						movement.Y = .1
						AxisMovement <- movement

					}
				}()
				if w.Keyboard().Up(keyboard.W) {
					movement.Z = 0

				}
				if w.Keyboard().Up(keyboard.S) {
					movement.Z = 0

				}

				if w.Keyboard().Up(keyboard.A) {
					rotation.Y = 0

				}
				if w.Keyboard().Up(keyboard.D) {
					rotation.Y = 0

				}

				if w.Keyboard().Up(keyboard.Q) {
					movement.X = 0

				}
				if w.Keyboard().Up(keyboard.E) {
					movement.X = 0

				}

				if w.Keyboard().Up(keyboard.X) {
					movement.Y = 0

				}
				if w.Keyboard().Up(keyboard.Z) {
					movement.Y = 0
				}

				AxisRotation <- rotation
				AxisMovement <- movement
				break

			case keyboard.Typed:
				log.Println(ev.String())
				switch ev.S {
				case "t":
					break
				}
			default:

			}
		default:
		}

		// Update the camera position, rotation etc (done in both windows).
		for _, object := range tripod.ObjectList {
			UpdateCamera(d, object.GetObject(), tripod.camera)

			// Draw the cubes as seen by the secondary camera.
			d.Draw(d.Bounds(), object.GetObject(), tripod.camera)
		}

		// Render the whole frame.
		d.Render()
	}
}
