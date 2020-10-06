package recursion

type  Point struct {
	x int
	y int
}

// Declare an enum to represent all possible movements from the current cell
type Movement int		// Used to declare an enum
const (
	NONE Movement = iota
	LEFT
	RIGHT
	TOP
	DOWN
)

func (m Movement) String() string {
	return [...]string{"None", "Left", "Right", "Top", "Down"}[m]
}


