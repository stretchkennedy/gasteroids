package obj

import (
	"log"
	"math"
	"math/rand"
	. "unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"

	"github.com/stretchkennedy/gasteroids/program"
)

type Asteroid struct {
	vao uint32
	vbo uint32
	program uint32
	Vertices []float32
}

const glAttrNum = 0
const glVecNum = 3

func NewAsteroid(sides int) *Asteroid {
	ast := &Asteroid{Vertices: make([]float32, (sides) * glVecNum)}

	for i := 0; i < sides; i++ {
		radiusModifier := rand.Float32() / 2.0 + 0.5
		angle := (math.Pi * 2.0) * (float64(i) / float64(sides))
		ast.Vertices[i*glVecNum + 0] = float32(math.Cos(angle)) * radiusModifier //x
		ast.Vertices[i*glVecNum + 1] = float32(math.Sin(angle)) * radiusModifier //y
	}

	ast.refreshGeometry()
	return ast
}

func (ast *Asteroid) Update(elapsed float64) {

}

func (ast *Asteroid) Render() {
	gl.UseProgram(ast.program)
	gl.BindVertexArray(ast.vao)
	gl.DrawArrays(gl.LINE_LOOP, glAttrNum, int32(len(ast.Vertices) / glVecNum))
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
	gl.BufferData(gl.ARRAY_BUFFER, len(ast.Vertices) * int(Sizeof(gl.FLOAT)), gl.Ptr(ast.Vertices), gl.STATIC_DRAW)

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

in vec3 vp;
void main() {
  gl_Position = vec4(vp, 1.0);
}
` + "\x00"

var fragmentShader string = `
#version 130

out vec4 frag_color;
void main() {
  frag_color = vec4(0.5, 0.0, 0.5, 1.0);
}
` + "\x00"
