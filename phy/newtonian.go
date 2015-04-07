package phy

import (
	. "github.com/go-gl/mathgl/mgl32"
)

type Newtonian struct {
	Position Vec2
	Velocity Vec2
	Rotation float32
}

func NewNewtonian(position, velocity Vec2, rotation float32) *Newtonian {
	return &Newtonian{
		Position: position,
		Velocity: velocity,
		Rotation: rotation,
	}
}

func (p *Newtonian) Update(elapsed float64) {
	p.Position =
		p.Position.Add(
		p.Velocity.Mul(
			float32(elapsed)))
}
