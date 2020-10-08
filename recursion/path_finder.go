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

func findPath(grid [][]int, entryPoint Point, exitPoint Point) (bool, []Point) {
	// Add necessary checks to ensure that entry and exit points are within
	// the given grid (omitted for brevity)
	// ...

	var path []Point
	hasSolution := solvePath(grid, entryPoint, exitPoint, &path);
	return hasSolution, path
}

func solvePath(grid [][]int, currentPoint Point, exitPoint Point, path *[]Point) bool {

}

func updateCurrentPoint(movement Movement, currentPoint Point) {
	switch movement {
		case LEFT: currentPoint.x--
		case RIGHT: currentPoint.x++
		case TOP: currentPoint.y--
		case DOWN: currentPoint.y++
	}
}

func resetPathPoint(movement Movement, currentPoint Point) {
	switch movement {
		case LEFT: currentPoint.x++
		case RIGHT: currentPoint.x--
		case TOP: currentPoint.y++
		case DOWN: currentPoint.y--
	}
}
