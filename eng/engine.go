package eng

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	. "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.1/glfw"


	"github.com/stretchkennedy/gasteroids/obj"
)

type Engine struct {
	Objects []obj.GameObject
	*glfw.Window
}

func init() {
	runtime.LockOSThread() 		           // switching GoRoutines between threads will break OpenGL
}

func NewEngine(winHeight, winWidth int) *Engine {
	window := NewWindow(winWidth, winHeight)
	window.MakeContextCurrent()
	return &Engine{Window: window}
}

func (eng *Engine) Start() {
	defer glfwTeardown()
	previousTime := glfw.GetTime()
	for !eng.Window.ShouldClose() {
		//// SETUP
		time := glfw.GetTime()
		elapsed := time - previousTime

		rawWidth, rawHeight:= eng.Window.GetFramebufferSize()
		height := float32(10.0)
		width := float32(rawWidth) / float32(rawHeight) * height

		projection := Ortho2D(0.0, width, height, 0.0) // 2d orthogonal, LRBT
		view := Ident4() // identity matrix
		vp := projection.Mul4(view)

		//// THINGS HAPPEN
		// clear buffer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// handle physics, controls, etc.
		for _, o := range eng.Objects {
			o.Update(height, width, elapsed)
		}

		// draw to window
		for _, o := range eng.Objects {
			o.Render(vp)
		}

		//// END
		previousTime = time
		eng.Window.SwapBuffers()
		glfw.PollEvents()
	}
}

func glSetup() {
	err := gl.Init()
	if err != nil {
		log.Panic(err)
	}

	gl.Enable(gl.DEPTH_TEST) // depth testing
	gl.DepthFunc(gl.LESS)    // smaller is closer
}

func glfwSetup()  {
	err := glfw.Init()
	if err != nil {
		log.Panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
}

func glfwTeardown() {
	glfw.Terminate()
}

func NewWindow(h int, w int) *glfw.Window {
	glSetup()
	glfwSetup()
	window, err := glfw.CreateWindow(h, w, "Gasteroids", glfw.GetPrimaryMonitor(), nil)
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	if err != nil {
		log.Panic(err)
	}
	return window
}
