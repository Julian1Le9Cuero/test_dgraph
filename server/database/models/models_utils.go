package models

// Interface to fill incurrent date
type SetDefault interface {
	setDate()
}

func FillDefault(s SetDefault) {
	s.setDate()
}
