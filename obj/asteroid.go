package obj

import (
	"math"
	"math/rand"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
	"github.com/stretchkennedy/gasteroids/phy"
)

type Asteroid struct {
	Radius float32
	Geometry geo.Geometry
	Physics *phy.Newtonian
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
		Physics: phy.NewNewtonian(position, velocity, 0),
		Geometry: geo.NewPolygon(vertices),
		Radius: maxRadius,
	}

	return ast
}

func (ast *Asteroid) Update(height, width float32, elapsed float64) {
	ast.Physics.Update(elapsed)

	d := ast.Radius * 2

	// for each dimension, wrap position
	if ast.Physics.Position[0] > width + d {
		ast.Physics.Position[0] = 0 - d
	}
	if ast.Physics.Position[0] < 0 - d {
		ast.Physics.Position[0] = width + d
	}
	if ast.Physics.Position[1] > height + d {
		ast.Physics.Position[1] = 0 - d
	}
	if ast.Physics.Position[1] < 0 - d {
		ast.Physics.Position[1] = height + d
	}
}

func (ast *Asteroid) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(ast.Physics.Position.X(), ast.Physics.Position.Y(), 0) // move model
	model = model.Mul4(HomogRotate3DZ(ast.Physics.Rotation))
	mvp := vp.Mul4(model)

	// render geometry
	ast.Geometry.Render(mvp)
}
