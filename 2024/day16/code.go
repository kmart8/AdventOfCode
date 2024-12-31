package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

type State struct {
	x    int
	y    int
	dir  int // 0=Up, 1=Right, 2=Down, 3=Left
	cost int
}

// PriorityQueue implements the heap.Interface and holds States.
type PriorityQueue []State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func valid(x int, y int, rows int, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func buildMatrix() ([][]string, Point, Point) {
	file, err := os.Open("./inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]string{}

	var start, end Point
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		rowSlice := []string{}
		for col, ch := range line {
			cell := string(ch)
			rowSlice = append(rowSlice, cell)

			if cell == "S" {
				start = Point{x: row, y: col}
			} else if cell == "E" {
				end = Point{x: row, y: col}
			}
		}
		matrix = append(matrix, rowSlice)
		row++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return matrix, start, end
}

func findLowestScore(grid [][]string, start Point, end Point) ([][][]int, int) {
	rows := len(grid)
	cols := len(grid[0])

	// dist[x][y][dir]: minimal cost to reach (x,y) with orientation=dir
	dist := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([][]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = make([]int, 4)
			for d := 0; d < 4; d++ {
				dist[i][j][d] = math.MaxInt64 // large number
			}
		}
	}

	// Offsets for directions: 0=Up, 1=Right, 2=Down, 3=Left
	directions := []Point{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Start facing East => dir=1
	dist[start.x][start.y][1] = 0
	heap.Push(pq, State{x: start.x, y: start.y, dir: 1, cost: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)
		cx, cy, cdir, ccost := current.x, current.y, current.dir, current.cost

		if cx == end.x && cy == end.y {
			return dist, ccost
		}

		// 1) Move forward
		fwdX := cx + directions[cdir].x
		fwdY := cy + directions[cdir].y
		if valid(fwdX, fwdY, rows, cols) && grid[fwdX][fwdY] != "#" {
			newCost := ccost + 1
			if newCost < dist[fwdX][fwdY][cdir] {
				dist[fwdX][fwdY][cdir] = newCost
				heap.Push(pq, State{x: fwdX, y: fwdY, dir: cdir, cost: newCost})
			}
		}

		// 2) Rotate left
		leftDir := (cdir + 3) % 4
		rotateLeftCost := ccost + 1000
		if rotateLeftCost < dist[cx][cy][leftDir] {
			dist[cx][cy][leftDir] = rotateLeftCost
			heap.Push(pq, State{x: cx, y: cy, dir: leftDir, cost: rotateLeftCost})
		}

		// 3) Rotate right
		rightDir := (cdir + 1) % 4
		rotateRightCost := ccost + 1000
		if rotateRightCost < dist[cx][cy][rightDir] {
			dist[cx][cy][rightDir] = rotateRightCost
			heap.Push(pq, State{x: cx, y: cy, dir: rightDir, cost: rotateRightCost})
		}
	}
	return dist, -1
}

// Predecessors of (x, y, dir) for reverse BFS
func predecessors(x, y, dir int, grid [][]string) []struct {
	nx, ny, ndir, stepCost int
} {
	var result []struct {
		nx, ny, ndir, stepCost int
	}

	// Directions: 0=Up, 1=Right, 2=Down, 3=Left
	offsets := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// If we reached (x,y,dir) by moving forward, original was (nx,ny,dir)
	dx, dy := offsets[dir].x, offsets[dir].y
	px := x - dx
	py := y - dy
	if valid(px, py, len(grid), len(grid[0])) && grid[px][py] != "#" {
		result = append(result, struct {
			nx, ny, ndir, stepCost int
		}{px, py, dir, 1})
	}

	// If we reached (x,y,dir) by rotating left, old dir = (dir+1)%4
	oldDirLeft := (dir + 1) % 4
	result = append(result, struct {
		nx, ny, ndir, stepCost int
	}{x, y, oldDirLeft, 1000})

	// If we reached (x,y,dir) by rotating right, old dir = (dir+3)%4
	oldDirRight := (dir + 3) % 4
	result = append(result, struct {
		nx, ny, ndir, stepCost int
	}{x, y, oldDirRight, 1000})

	return result
}

// backtrackBestPathCells returns the set of (x,y) cells on at least one best path
func backtrackBestPathCells(grid [][]string, dist [][][]int, end Point) map[Point]bool {
	rows := len(grid)
	cols := len(grid[0])

	// Find minimal cost at end
	bestCost := math.MaxInt64
	for d := 0; d < 4; d++ {
		if dist[end.x][end.y][d] < bestCost {
			bestCost = dist[end.x][end.y][d]
		}
	}
	if bestCost == math.MaxInt64 {
		// no path
		return nil
	}

	// Mark visited states for reverse BFS
	visited := make([][][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([][]bool, cols)
		for j := 0; j < cols; j++ {
			visited[i][j] = make([]bool, 4)
		}
	}

	// Start from all end dirs that achieve bestCost
	queue := []State{}
	for d := 0; d < 4; d++ {
		if dist[end.x][end.y][d] == bestCost {
			visited[end.x][end.y][d] = true
			queue = append(queue, State{x: end.x, y: end.y, dir: d})
		}
	}

	// Reverse BFS
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cx, cy, cdir := cur.x, cur.y, cur.dir
		curDist := dist[cx][cy][cdir]

		for _, p := range predecessors(cx, cy, cdir, grid) {
			nx, ny, ndir, stepCost := p.nx, p.ny, p.ndir, p.stepCost
			if dist[nx][ny][ndir]+stepCost == curDist {
				if !visited[nx][ny][ndir] {
					visited[nx][ny][ndir] = true
					queue = append(queue, State{x: nx, y: ny, dir: ndir})
				}
			}
		}
	}

	// Gather cells that appear in visited states
	bestCells := make(map[Point]bool)
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for d := 0; d < 4; d++ {
				if visited[x][y][d] {
					bestCells[Point{x, y}] = true
					break
				}
			}
		}
	}
	return bestCells
}

func main() {
	grid, start, end := buildMatrix()
	dist, bestCost := findLowestScore(grid, start, end)
	bestCells := backtrackBestPathCells(grid, dist, end)
	fmt.Println(fmt.Sprintf("The answer to day 16 part 1 is %d", bestCost))
	fmt.Println(fmt.Sprintf("The answer to day 16 part 2 is %d", len(bestCells)))
}
