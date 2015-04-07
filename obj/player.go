package obj

import (
	"math"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
)

type Player struct {
	Position Vec2
	Velocity Vec2
	Rotation float32

	radius float32
	geometry geo.Geometry
}

func NewPlayer(position Vec2) (pl *Player) {
	sides := 3
	vertices := make([]float32, (sides) * glVecNum)
	for i := 0; i < sides; i++ {
		angle := (math.Pi * 2.0) * (float64(i) / float64(sides))
		vertices[i*glVecNum + 0] = float32(math.Cos(angle)) //x
		vertices[i*glVecNum + 1] = float32(math.Sin(angle)) //y
	}

	pl = &Player{
		Position: position,
		geometry: geo.NewPolygon(vertices),
		radius: 1,
	}

	return pl
}

func (pl *Player) Update(height, width float32, elapsed float64) {
	pl.Rotation += float32(elapsed)
	pl.Position =
		pl.Position.Add(
		pl.Velocity.Mul(
		float32(elapsed)))
	d := pl.radius * 2

	// for each dimension, wrap position
	if pl.Position[0] > width + d {
		pl.Position[0] = 0 - d
	}
	if pl.Position[0] < 0 - d {
		pl.Position[0] = width + d
	}
	if pl.Position[1] > height + d {
		pl.Position[1] = 0 - d
	}
	if pl.Position[1] < 0 - d {
		pl.Position[1] = height + d
	}
}

func (pl *Player) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(pl.Position.X(), pl.Position.Y(), 0) // move model
	model = model.Mul4(HomogRotate3DZ(pl.Rotation))
	mvp := vp.Mul4(model)

	// render geometry
	pl.geometry.Render(mvp)
}
