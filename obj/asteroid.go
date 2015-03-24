package obj

import (
	"log"
	. "unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"

	"github.com/stretchkennedy/gasteroids/program"
)

type GameObject interface {
	Render()
	Update(delta int)
}

type Asteroid struct {
	vao uint32
	vbo uint32
	program uint32
	Vertices []float32
}

func NewAsteroid(vertices []float32) *Asteroid {
	ast := &Asteroid{Vertices: vertices}
	ast.setup()
	return ast
}

const glAttrNum = 0
const glVecNum = 3

func (ast *Asteroid) setup() {
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

func (ast *Asteroid) Render() {
	gl.UseProgram(ast.program)
	gl.BindVertexArray(ast.vao)
	gl.DrawArrays(gl.TRIANGLES, glAttrNum, glVecNum)
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
