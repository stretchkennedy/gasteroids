package obj

import (
	"math"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
	"github.com/stretchkennedy/gasteroids/phy"
)

type Player struct {
	Radius float32
	Geometry geo.Geometry
	Physics *phy.Newtonian
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
		Physics: phy.NewNewtonian(position, Vec2{0,0}, 0, 1),
		Geometry: geo.NewPolygon(vertices),
	}

	return pl
}

func (pl *Player) Update(height, width float32, elapsed float64) {
	pl.Physics.Rotation += float32(elapsed)
	pl.Physics.Update(height, width, elapsed)
}

func (pl *Player) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(pl.Physics.Position.X(), pl.Physics.Position.Y(), 0) // move model
	model = model.Mul4(HomogRotate3DZ(pl.Physics.Rotation))
	mvp := vp.Mul4(model)

	// render geometry
	pl.Geometry.Render(mvp)
}
