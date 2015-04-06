package obj

import (
	"math"
	"math/rand"

	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/geo"
)

type Asteroid struct {
	radius float32
	position Vec2
	velocity Vec2
	geometry geo.Geometry
}

const glVecNum = 3

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
		geometry: geo.NewPolygon(vertices),
		position: position,
		velocity: velocity,
		radius: maxRadius,
	}

	return ast
}

func (ast *Asteroid) Update(height, width float32, elapsed float64) {
	ast.position =
		ast.position.Add(
		ast.velocity.Mul(
			float32(elapsed)))

	// for each dimension, wrap position
	if ast.position[0] > width + ast.radius {
		ast.position[0] = 0 - ast.radius
	}
	if ast.position[0] < 0 - ast.radius {
		ast.position[0] = width + ast.radius
	}
	if ast.position[1] > height + ast.radius {
		ast.position[1] = 0 - ast.radius
	}
	if ast.position[1] < 0 - ast.radius {
		ast.position[1] = height + ast.radius
	}
}

func (ast *Asteroid) Render(vp Mat4) {
	// MVP matrices
	model := Translate3D(ast.position.X(), ast.position.Y(), 0) // move model
	mvp := vp.Mul4(model)

	// render geometry
	ast.geometry.Render(mvp)
}
