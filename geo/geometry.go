package geo

import (
	. "github.com/go-gl/mathgl/mgl32"
)

type Geometry interface {
	Render(mvp Mat4)
}
