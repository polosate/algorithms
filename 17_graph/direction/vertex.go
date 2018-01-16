package non_direction

type vertex struct {
	label      string
	wasVisited bool
}

func newVertex(label string) vertex {
	return vertex{
		label:      label,
		wasVisited: false,
	}
}
