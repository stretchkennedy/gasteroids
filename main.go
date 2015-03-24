package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/stretchkennedy/gasteroids/obj"
)

const HEIGHT = 640
const WIDTH = 480

func init() {
	runtime.LockOSThread()
}

func main() {
	//// SETUP
	// gl/glfw setup
	glSetup()
	glfwSetup()
	defer glfwTeardown()

	// window setup
	window := NewWindow(HEIGHT, WIDTH)
	window.MakeContextCurrent()

	// game setup
	verts := []float32{
		0.0, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}
	ast := obj.NewAsteroid(verts)

	//// MAIN LOOP
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		ast.Render()

		window.SwapBuffers()
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
}

func glfwTeardown() {
	glfw.Terminate()
}

func NewWindow(h int, w int) *glfw.Window {
	window, err := glfw.CreateWindow(h, w, "Gasteroids", nil, nil)
	if err != nil {
		log.Panic(err)
	}
	return window
}
