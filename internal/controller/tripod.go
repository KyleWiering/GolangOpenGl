package controller

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/camera"
	"log"
	"time"

	//"azul3d.org/engine/gfx/gfxutil"
	"azul3d.org/engine/lmath"
	//"github.com/waryway/go2dgame/internal/models"
	//"image"
	//"log"
	"math"
)

type Tripod struct {
	device      gfx.Device
	camera      *camera.Camera
	coordinates Coordinates
	ObjectList  []Object
}

func (t Tripod) GetCamera() *camera.Camera {
	return t.camera
}

func (t *Tripod) Movement(AxisMovement chan (lmath.Vec3)) {

	for {
		time.Sleep(time.Microsecond * 10)
		select {
		case axis := <-AxisMovement:
			t.coordinates.MoveXAxis(axis.X)
			t.coordinates.MoveYAxis(axis.Y)
			t.coordinates.MoveZAxis(axis.Z)
			log.Println(axis, t.coordinates, t.camera.Pos())

		}
		t.camera.SetPos(t.coordinates.ObjectPosition)
	}
}
func (t *Tripod) Rotation(RotationMovement chan (lmath.Vec3)) {
	for {
		time.Sleep(time.Microsecond * 10)
		select {
		case rotation := <-RotationMovement:
			t.coordinates.RotateX(rotation.X)
			t.coordinates.RotateY(rotation.Y)
			t.coordinates.RotateZ(rotation.Z)

			t.camera.SetRot(lmath.Vec3{t.coordinates.xRotation, t.coordinates.yRotation, t.coordinates.zRotation})
		}
	}
}

func (t *Tripod) Init(d gfx.Device, objects []Object) {
	t.device = d
	t.ObjectList = objects
	t.coordinates.Initialize()
	// Create and position both cameras.
	t.camera = camera.New(t.device.Bounds())
	t.camera.SetPos(t.coordinates.ObjectPosition)
	t.camera.SetRot(lmath.Vec3{t.coordinates.xRotation, t.coordinates.yRotation, t.coordinates.zRotation})

}

func UpdateCamera(d gfx.Device, cubes *gfx.Object, camMain *camera.Camera) {
	// Clear the entire area.
	d.Clear(d.Bounds(), gfx.Color{0, 0, 0, 1})
	d.ClearDepth(d.Bounds(), 1.0)

	// Use frame delta time so the windows stay in sync.
	//dt := d.Clock().Dt()
	//cubes.SetRot(cubes.Rot().AddScalar(10 * dt))

	// How much we want to add to the FOV and far clipping distance (pulsating).
	add := math.Sin(float64(d.Clock().FrameCount())/20) * 25

	camMain.Near = 1
	camMain.Far = 30 + add

	// Set camera perspective and UpdateCamera the position of the camera in case
	// we were previously using an orthographic camera.
	camMain.FOV = 75 + add
	camMain.Update(d.Bounds())
	//	camMain.SetPos(lmath.Vec3{0, -10, 25})
}
