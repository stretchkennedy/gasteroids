package obj

type GameObject interface {
	Render()
	Update(elapsed float64)
}
