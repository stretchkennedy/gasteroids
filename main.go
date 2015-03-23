package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

const HEIGHT = 640
const WIDTH = 480
const WINDOW_NAME = "Gasteroids"

func init() {
	runtime.LockOSThread()
}

func main() {
	// setup
	err := gl.Init()
	if err != nil {
		log.Panic(err)
	}

	err = glfw.Init()
	if err != nil {
		log.Panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(HEIGHT, WIDTH, WINDOW_NAME, nil, nil)
	if err != nil {
		log.Panic(err)
	}
	window.MakeContextCurrent()

	// main loop
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
