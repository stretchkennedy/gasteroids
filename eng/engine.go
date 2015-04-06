package eng

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	. "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.1/glfw"


	"github.com/stretchkennedy/gasteroids/obj"
)

type Engine struct {
	objects []obj.GameObject
	window *glfw.Window
}

func NewEngine(window *glfw.Window, objects []obj.GameObject) *Engine {
	return &Engine{objects: objects, window: window}
}

func (eng *Engine) Start() {
	previousTime := glfw.GetTime()
	for !eng.window.ShouldClose() {
		//// SETUP
		time := glfw.GetTime()
		elapsed := time - previousTime

		rawWidth, rawHeight:= eng.window.GetFramebufferSize()
		height := float32(10.0)
		width := float32(rawWidth) / float32(rawHeight) * height

		projection := Ortho2D(0.0, width, height, 0.0) // 2d orthogonal, LRBT
		view := Ident4() // identity matrix
		vp := projection.Mul4(view)

		//// THINGS HAPPEN
		// clear buffer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// handle physics, controls, etc.
		for _, o := range eng.objects {
			o.Update(height, width, elapsed)
		}

		// draw to window
		for _, o := range eng.objects {
			o.Render(vp)
		}

		//// END
		previousTime = time
		eng.window.SwapBuffers()
		glfw.PollEvents()
	}
}
