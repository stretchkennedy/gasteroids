package obj

import (
	. "github.com/go-gl/mathgl/mgl32"
)

type GameObject interface {
	Render(vp Mat4)
	Update(height, width float32, elapsed float64)
}
