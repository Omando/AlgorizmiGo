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

// Can only move left, top, up and down. No diagonal movement
func getMovements(grid [][]int, currentPoint Point) []Movement {
	var movements []Movement
	if currentPoint.y+1 < len(grid) && grid[currentPoint.y + 1][currentPoint.x] == 1 {
		movements = append(movements, DOWN);
	}
	if currentPoint.y-1 >= 0 && grid[currentPoint.y - 1][currentPoint.x] == 1 {
		movements = append(movements, TOP);
	}
	if currentPoint.x-1 >= 0 && grid[currentPoint.y][currentPoint.x - 1] == 1 {
		movements = append(movements, LEFT);
	}
	if currentPoint.x+1 < len(grid[0]) && grid[currentPoint.y][currentPoint.x + 1] == 1 {
		movements = append(movements, RIGHT);
	}
	if len(movements) == 0 {
		movements = append(movements, NONE);
	}

	return movements;
}

func updateCurrentPoint(movement Movement, currentPoint Point) {
	switch movement {
		case LEFT:
			currentPoint.x--
		case RIGHT:
			currentPoint.x++
		case TOP:
			currentPoint.y--
		case DOWN:
			currentPoint.y++
	}
}

func resetPathPoint(movement Movement, currentPoint Point) {
	switch movement {
		case LEFT:
			currentPoint.x++
		case RIGHT:
			currentPoint.x--
		case TOP:
			currentPoint.y++
		case DOWN:
			currentPoint.y--
	}
}
