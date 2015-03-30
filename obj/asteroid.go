package obj

import (
	"log"
	"math"
	"math/rand"
	. "unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/program"
)

type Asteroid struct {
	r float32
	position Vec2
	velocity Vec2

	vao uint32
	vbo uint32
	program uint32
	vertices []float32
}

const glAttrNum = 0
const glVecNum = 3

func NewAsteroid(sides int, position, velocity Vec2) *Asteroid {
	ast := &Asteroid{
		vertices: make([]float32, (sides) * glVecNum),
		position: position,
		velocity: velocity,
	}

	for i := 0; i < sides; i++ {
		radiusModifier := rand.Float32() / 2.0 + 0.5
		angle := (math.Pi * 2.0) * (float64(i) / float64(sides))
		ast.vertices[i*glVecNum + 0] = float32(math.Cos(angle)) * radiusModifier //x
		ast.vertices[i*glVecNum + 1] = float32(math.Sin(angle)) * radiusModifier //y
		ast.r = float32(math.Max(float64(radiusModifier), float64(ast.r))) //circle describing outer bounds
	}

	ast.refreshGeometry()
	return ast
}

func (ast *Asteroid) Update(height, width float32, elapsed float64) {
	ast.position =
		ast.position.Add(
		ast.velocity.Mul(
			float32(elapsed)))

	// for each dimension, wrap position
	if ast.position[0] > width + ast.r {
		ast.position[0] = 0 - ast.r
	}
	if ast.position[0] < 0 - ast.r {
		ast.position[0] = width + ast.r
	}
	if ast.position[1] > height + ast.r {
		ast.position[1] = 0 - ast.r
	}
	if ast.position[1] < 0 - ast.r {
		ast.position[1] = height + ast.r
	}
}

func (ast *Asteroid) Render(height, width float32) {
	//// MVP matrices
    // setup projection
	projection := Ortho2D(0.0, width, height, 0.0) // 2d orthogonal, LRBT
	projectionLoc := gl.GetUniformLocation(ast.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])

	// setup camera
	view := Ident4() // identity matrix
	viewLoc := gl.GetUniformLocation(ast.program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])

	// setup model
	model := Translate3D(ast.position.X(), ast.position.Y(), 0) // move model
	modelLoc := gl.GetUniformLocation(ast.program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

	//// load relevant things
	gl.UseProgram(ast.program)
	gl.BindVertexArray(ast.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, ast.vbo)

	//// draw geometry
	gl.DrawArrays(gl.LINE_LOOP, glAttrNum, int32(len(ast.vertices) / glVecNum))
}

func (ast *Asteroid) refreshGeometry() {
	// remove previous geometry data
	ast.clearGeometry()

	// setup vao
	gl.GenVertexArrays(1, &ast.vao)
	gl.BindVertexArray(ast.vao)

	// setup vbo
	gl.GenBuffers(1, &ast.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, ast.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(ast.vertices) * int(Sizeof(gl.FLOAT)), gl.Ptr(ast.vertices), gl.STATIC_DRAW)

	// setup attribute array
	gl.EnableVertexAttribArray(glAttrNum)
	gl.VertexAttribPointer(glAttrNum, glVecNum, gl.FLOAT, false, 0, nil)

	// setup shaders
	var err error
	ast.program, err = program.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		log.Panic(err)
	}
}

func (ast *Asteroid) clearGeometry() {
	gl.DeleteVertexArrays(1, &ast.vao)
	gl.DeleteBuffers(1, &ast.vbo)
	gl.DeleteProgram(ast.program)
}

var vertexShader string = `
#version 130

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

in vec3 vp;

void main() {
  gl_Position = projection * view * model * vec4(vp, 1.0);
}
` + "\x00"

var fragmentShader string = `
#version 130

out vec4 frag_color;
void main() {
  frag_color = vec4(0.5, 0.0, 0.5, 1.0);
}
` + "\x00"
