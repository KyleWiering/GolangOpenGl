package models

import "azul3d.org/engine/gfx"

// cube returns a cube *gfx.Mesh at an offset from the origin
func Cube(x, y float32) *gfx.Mesh {
	m := gfx.NewMesh()
	m.Vertices = []gfx.Vec3{
		{-0.5 + x, -0.5 + y, 0.5},
		{0.5 + x, -0.5 + y, 0.5},
		{0.5 + x, 0.5 + y, 0.5},
		{-0.5 + x, 0.5 + y, 0.5},

		{-0.5 + x, -0.5 + y, -0.5},
		{-0.5 + x, 0.5 + y, -0.5},
		{0.5 + x, 0.5 + y, -0.5},
		{0.5 + x, -0.5 + y, -0.5},

		{-0.5 + x, 0.5 + y, -0.5},
		{-0.5 + x, 0.5 + y, 0.5},
		{0.5 + x, 0.5 + y, 0.5},
		{0.5 + x, 0.5 + y, -0.5},

		{-0.5 + x, -0.5 + y, -0.5},
		{0.5 + x, -0.5 + y, -0.5},
		{0.5 + x, -0.5 + y, 0.5},
		{-0.5 + x, -0.5 + y, 0.5},

		{0.5 + x, -0.5 + y, -0.5},
		{0.5 + x, 0.5 + y, -0.5},
		{0.5 + x, 0.5 + y, 0.5},
		{0.5 + x, -0.5 + y, 0.5},

		{-0.5 + x, -0.5 + y, -0.5},
		{-0.5 + x, -0.5 + y, 0.5},
		{-0.5 + x, 0.5 + y, 0.5},
		{-0.5 + x, 0.5 + y, -0.5},
	}

	m.Indices = []uint32{
		0, 1, 2, 0, 2, 3, // front
		4, 5, 6, 4, 6, 7, // back
		8, 9, 10, 8, 10, 11, // top
		12, 13, 14, 12, 14, 15, // bottom
		16, 17, 18, 16, 18, 19, // right
		20, 21, 22, 20, 22, 23, // left
	}

	return m
}
