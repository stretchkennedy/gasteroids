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
	"github.com/stretchkennedy/gasteroids/eng"
)

func init() {
	runtime.LockOSThread() 		           // switching GoRoutines between threads will break OpenGL
	rand.Seed(time.Now().UTC().UnixNano()) // make random behave unpredictably
}

func main() {
	//// PARSE OPTIONS
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
	window := NewWindow(winWidth, winHeight)
	window.MakeContextCurrent()

	// game setup
	player := obj.NewPlayer(Vec2{5,5})
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			switch key {
			case glfw.KeyLeft:
				player.Physics.RotationalVelocity = -1
			case glfw.KeyRight:
				player.Physics.RotationalVelocity = 1
			}
		case glfw.Release:
			switch key {
			case glfw.KeyLeft:
				player.Physics.RotationalVelocity = 0
			case glfw.KeyRight:
				player.Physics.RotationalVelocity = 0
			}
		}
	})

	objects := []obj.GameObject{
		obj.NewAsteroid(12, Vec2{3, 3}, Vec2{2, 1}),
		obj.NewAsteroid(5, Vec2{7, 7}, Vec2{1, 2}),
		player,
	}
	engine := eng.NewEngine(window, objects)

	//// MAIN LOOP
	engine.Start()
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
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	if err != nil {
		log.Panic(err)
	}
	return window
}
