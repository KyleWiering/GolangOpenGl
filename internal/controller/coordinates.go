package controller

import (
	"azul3d.org/engine/lmath"
	"math"
)

// Coordinates are an object's 3d space position, allowing for movement along axis, inluding rotations.
type Coordinates struct {
	ObjectPosition lmath.Vec3
	xLocation      lmath.Vec3
	yLocation      lmath.Vec3
	zLocation      lmath.Vec3
	xRotation      float64
	yRotation      float64
	zRotation      float64
}

func (c *Coordinates) Initialize() {
	c.ObjectPosition = lmath.Vec3{0, 0, 0}
	c.xLocation = lmath.Vec3{0, 0, -1}
	c.yLocation = lmath.Vec3{1, 0, 0}
	c.zLocation = lmath.Vec3{0, 1, 0}

	c.xRotation = 0.0
	c.yRotation = 0.0
	c.zRotation = 0.0
}

// RotateX Rotate along the x axis the provided angle
func (c *Coordinates) RotateX(angle float64) {
	c.xRotation += angle

	c.zLocation = c.zLocation.MulScalar(math.Cos(lmath.Radians(angle))).Add(c.yLocation.MulScalar(math.Sin(lmath.Radians(angle))))
	normal, ok := c.zLocation.Normalized()
	if ok {
		c.zLocation = normal
	}

	c.yLocation = c.zLocation.Cross(c.xLocation).MulScalar(-1)
}

// RotateY Rotate along the y axis the provided angle
func (c *Coordinates) RotateY(angle float64) {
	c.yRotation += angle

	c.zLocation = c.zLocation.MulScalar(math.Cos(lmath.Radians(angle))).Add(c.xLocation.MulScalar(math.Sin(lmath.Radians(angle))))
	normal, ok := c.zLocation.Normalized()
	if ok {
		c.zLocation = normal
	}

	c.xLocation = c.zLocation.Cross(c.yLocation)
}

// RotateZ Rotate along the Z axis the provided angle
func (c *Coordinates) RotateZ(angle float64) {
	c.zRotation += angle

	c.xLocation = c.xLocation.MulScalar(math.Cos(lmath.Radians(angle))).Add(c.yLocation.MulScalar(math.Sin(lmath.Radians(angle))))
	normal, ok := c.xLocation.Normalized()
	if ok {
		c.xLocation = normal
	}

	c.yLocation = c.zLocation.Cross(c.xLocation).MulScalar(-1)
}

// MoveXAxis Move the object along the x axis.
func (c *Coordinates) MoveXAxis(distance float64) {
	c.ObjectPosition = c.ObjectPosition.Add(c.xLocation.MulScalar(distance))
}

// MoveYAxis Move the object along the y axis.
func (c *Coordinates) MoveYAxis(distance float64) {
	c.ObjectPosition = c.ObjectPosition.Add(c.yLocation.MulScalar(distance))
}

// MoveZAxis Move the object along the Z axis.
func (c *Coordinates) MoveZAxis(distance float64) {
	c.ObjectPosition = c.ObjectPosition.Add(c.zLocation.MulScalar(-distance))
}
