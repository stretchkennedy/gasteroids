package obj

type GameObject interface {
	Render(height, width float32)
	Update(height, width float32, elapsed float64)
}
