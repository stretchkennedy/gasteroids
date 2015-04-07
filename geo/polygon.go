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

func (poly *Polygon) Render(mvp Mat4) {
	// load relevant things
	gl.UseProgram(poly.program)
	gl.BindVertexArray(poly.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, poly.vbo)

	// MVP matrix
	mvpLoc := gl.GetUniformLocation(poly.program, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvpLoc, 1, false, &mvp[0])

	// draw geometry
	gl.DrawArrays(gl.LINE_LOOP, glAttrNum, int32(len(poly.vertices) / glVecNum))
}

var vertexShader string = `
#version 130

uniform mat4 mvp;

in vec3 vert;

void main() {
  gl_Position = mvp * vec4(vert, 1.0);
}
` + "\x00"

var fragmentShader string = `
#version 130

out vec4 frag_color;
void main() {
  frag_color = vec4(0.5, 0.0, 0.5, 1.0);
}
` + "\x00"
