package phy

import (
	. "github.com/go-gl/mathgl/mgl32"
)

type Newtonian struct {
	Radius float32
	Position Vec2
	Velocity Vec2
	Rotation float32
}

func NewNewtonian(position, velocity Vec2, rotation, radius float32) *Newtonian {
	return &Newtonian{
		Position: position,
		Velocity: velocity,
		Rotation: rotation,
		Radius: radius,
	}
}

func (p *Newtonian) Update(height, width float32, elapsed float64) {
	p.Position =
		p.Position.Add(
		p.Velocity.Mul(
			float32(elapsed)))


	d := p.Radius * 2

	// for each dimension, wrap position
	if p.Position[0] > width + d {
		p.Position[0] = 0 - d
	}
	if p.Position[0] < 0 - d {
		p.Position[0] = width + d
	}
	if p.Position[1] > height + d {
		p.Position[1] = 0 - d
	}
	if p.Position[1] < 0 - d {
		p.Position[1] = height + d
	}
}
