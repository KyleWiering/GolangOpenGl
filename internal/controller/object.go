package controller

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/gfxutil"
	"github.com/waryway/go2dgame/internal/models"
	"log"

	"azul3d.org/examples/abs"
)

type Object struct {
	coordinates Coordinates
	object      *gfx.Object
}

func (o *Object) GetObject() *gfx.Object {
	return o.object
}

func (o *Object) Init(objectName string) {
	// Load the shader.
	o.coordinates.Initialize()
	shader, err := gfxutil.OpenShader(abs.Path(objectName))
	if err != nil {
		log.Fatal(err)
	}
	// Insert some cubes so we have something to look at.
	o.object = gfx.NewObject()
	o.object.State = gfx.NewState()
	o.object.State.FaceCulling = gfx.BackFaceCulling
	o.object.Shader = shader
	o.object.Meshes = []*gfx.Mesh{}
	for x := -3; x <= 3; x++ {
		for y := -3; y <= 3; y++ {
			o.object.Meshes = append(o.object.Meshes, models.Cube(float32(x)*1.5, float32(y)*1.5))
		}
	}
}
