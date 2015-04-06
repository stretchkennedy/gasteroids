package geo

import (
	"log"
	. "unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	. "github.com/go-gl/mathgl/mgl32"

	"github.com/stretchkennedy/gasteroids/program"
)

type Polygon struct {
	vao uint32
	vbo uint32
	program uint32
	vertices []float32
}

type Geometry interface {
	Render(p Mat4, v Mat4, m Mat4)
}

const glAttrNum = 0
const glVecNum = 3

func NewPolygon(vertices []float32) (poly *Polygon) {
	poly = &Polygon{
		vertices: vertices,
	}
	poly.refresh()
	return poly
}

func (poly *Polygon) clear() {
	gl.DeleteVertexArrays(1, &poly.vao)
	gl.DeleteBuffers(1, &poly.vbo)
	gl.DeleteProgram(poly.program)
}

func (poly *Polygon) refresh() {
	// remove previous geometry data
	poly.clear()

	// setup vao
	gl.GenVertexArrays(1, &poly.vao)
	gl.BindVertexArray(poly.vao)

	// setup vbo
	gl.GenBuffers(1, &poly.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, poly.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(poly.vertices) * int(Sizeof(gl.FLOAT)), gl.Ptr(poly.vertices), gl.STATIC_DRAW)

	// setup attribute array
	gl.EnableVertexAttribArray(glAttrNum)
	gl.VertexAttribPointer(glAttrNum, glVecNum, gl.FLOAT, false, 0, nil)

	// setup shaders
	var err error
	poly.program, err = program.NewProgram(vertexShader, fragmentShader)
	if err != nil {
		log.Panic(err)
	}
}

func (poly *Polygon) Render(p Mat4, v Mat4, m Mat4) {
	// TODO: calculate matrix at each step and pass it down
	//// MVP matrices
    // setup projection
	projectionLoc := gl.GetUniformLocation(poly.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionLoc, 1, false, &p[0])

	// setup camera
	viewLoc := gl.GetUniformLocation(poly.program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewLoc, 1, false, &v[0])

	// setup model
	modelLoc := gl.GetUniformLocation(poly.program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelLoc, 1, false, &m[0])

	//// load relevant things
	gl.UseProgram(poly.program)
	gl.BindVertexArray(poly.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, poly.vbo)

	//// draw geometry
	gl.DrawArrays(gl.LINE_LOOP, glAttrNum, int32(len(poly.vertices) / glVecNum))
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
