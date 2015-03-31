package main

import (
	"log"
	"runtime"
	"math/rand"
	"time"
	"os"
	"strconv"

	"github.com/go-gl/gl/v4.5-core/gl"
	. "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/stretchkennedy/gasteroids/obj"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	winHeight, err := strconv.Atoi(os.Getenv("HEIGHT"))
	if err != nil {
		winHeight = 480
	}
	winWidth, err := strconv.Atoi(os.Getenv("WIDTH"))
	if err != nil {
		winWidth = 640
	}

	//// SETUP
	// gl/glfw setup
	glSetup()
	glfwSetup()
	defer glfwTeardown()

	// window setup
	window := NewWindow(winWidth, winHeight)
	window.MakeContextCurrent()

	// game setup
	rand.Seed(time.Now().UTC().UnixNano())
	var objects []obj.GameObject
	objects = []obj.GameObject{obj.NewAsteroid(9, Vec2{3, 3},  Vec2{2, 1}), obj.NewAsteroid(9, Vec2{7, 7}, Vec2{1, 2})}

	//// MAIN LOOP
	previousTime := glfw.GetTime()
	for !window.ShouldClose() {
		// retrieve information
		time := glfw.GetTime()
		elapsed := time - previousTime

		rawWidth, rawHeight:= window.GetFramebufferSize()
		height := float32(10.0)
		width := float32(rawWidth) / float32(rawHeight) * height

		// clear buffer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// handle physics, controls, etc.
		for _, o := range objects {
			o.Update(height, width, elapsed)
		}

		// draw to window
		for _, o := range objects {
			o.Render(height, width)
		}

		// end
		previousTime = time
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
	glfw.WindowHint(glfw.Resizable, glfw.False)
}

func glfwTeardown() {
	glfw.Terminate()
}

func NewWindow(h int, w int) *glfw.Window {
	window, err := glfw.CreateWindow(h, w, "Gasteroids", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		log.Panic(err)
	}
	return window
}
