package obj

import (
	"math"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
)

type Player struct {
	radius float32
	position Vec2
	rotation float32
	velocity Vec2
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
		geometry: geo.NewPolygon(vertices),
		position: position,
		radius: 1,
	}

	return pl
}

func (pl *Player) Update(height, width float32, elapsed float64) {
	pl.rotation += float32(elapsed)
	pl.position =
		pl.position.Add(
		pl.velocity.Mul(
		float32(elapsed)))
	d := pl.radius * 2

	// for each dimension, wrap position
	if pl.position[0] > width + d {
		pl.position[0] = 0 - d
	}
	if pl.position[0] < 0 - d {
		pl.position[0] = width + d
	}
	if pl.position[1] > height + d {
		pl.position[1] = 0 - d
	}
	if pl.position[1] < 0 - d {
		pl.position[1] = height + d
	}
}

func (pl *Player) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(pl.position.X(), pl.position.Y(), 0) // move model
	model = model.Mul4(HomogRotate3DZ(pl.rotation))
	mvp := vp.Mul4(model)

	// render geometry
	pl.geometry.Render(mvp)
}
