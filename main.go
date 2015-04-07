package main

import (
	"math/rand"
	"time"
	"os"
	"strconv"

	. "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/stretchkennedy/gasteroids/obj"
	"github.com/stretchkennedy/gasteroids/eng"
)

func init() {
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

	engine := eng.NewEngine(winHeight, winWidth)

	//// game setup
	player := obj.NewPlayer(Vec2{5,5})
	engine.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
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

	engine.Objects = []obj.GameObject{
		obj.NewAsteroid(12, Vec2{3, 3}, Vec2{2, 1}),
		obj.NewAsteroid(5, Vec2{7, 7}, Vec2{1, 2}),
		player,
	}

	//// MAIN LOOP
	engine.Start()
}
