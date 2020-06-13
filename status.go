package health

type Status string

const (
	StatusUp   = Status("UP")
	StatusDown = Status("DOWN")
)
