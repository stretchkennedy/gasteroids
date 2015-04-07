package obj

import (
	"math"
	"math/rand"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
	"github.com/stretchkennedy/gasteroids/phy"
)

type Asteroid struct {
	Geometry geo.Geometry
	Physics *phy.Wrapping
}

func NewAsteroid(sides int, position, velocity Vec2) (ast *Asteroid) {
	vertices := make([]float32, (sides) * glVecNum)
	maxRadius := float32(0.0)
	for i := 0; i < sides; i++ {
		radiusModifier := (rand.Float32() / 2.0 + 0.5) * (float32(sides) / 10.0)
		angle := (math.Pi * 2.0) * (float64(i) / float64(sides))
		vertices[i*glVecNum + 0] = float32(math.Cos(angle)) * radiusModifier //x
		vertices[i*glVecNum + 1] = float32(math.Sin(angle)) * radiusModifier //y
		maxRadius = float32(math.Max(float64(radiusModifier), float64(maxRadius))) //circle describing outer bounds
	}

	ast = &Asteroid{
		Physics: &phy.Wrapping{
			Position: position,
			Velocity: velocity,
			RotationalVelocity: 0.2,
			Radius: maxRadius,
		},
		Geometry: geo.NewPolygon(vertices),
	}

	return ast
}

func (ast *Asteroid) Update(height, width float32, elapsed float64) {
	ast.Physics.Update(height, width, elapsed)
}

func (ast *Asteroid) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(ast.Physics.Position.X(), ast.Physics.Position.Y(), 0) // move model
	model = model.Mul4(HomogRotate3DZ(ast.Physics.Rotation))
	mvp := vp.Mul4(model)

	// render geometry
	ast.Geometry.Render(mvp)
}
