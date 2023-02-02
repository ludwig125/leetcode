package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	tests := map[string]struct {
		grid [][]int
		want int
	}{
		"1": {
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
		"2": {
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
		"3": {
			grid: [][]int{
				{1, 1, 1},
				{0, 0, 0},
				{0, 1, 1},
			},
			want: 3,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maxAreaOfIsland(tc.grid)

			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestDfs(t *testing.T) {
	tests := map[string]struct {
		grid  [][]int
		input []int
		want  [][]int
	}{
		"1": {
			grid: [][]int{
				{1, 1, 0, 1, 0},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			input: []int{0, 0},
			want: [][]int{
				{0, 0},
				{1, 0},
				{0, 1},
				{0, 2},
				{1, 2},
			},
		},
		"2": {
			grid: [][]int{
				{1, 1, 0, 1, 0},
				{1, 0, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			input: []int{3, 0},
			want: [][]int{
				{3, 0},
				{3, 1},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			x, y := tc.input[0], tc.input[1]
			var visited [][]int
			visited = dfs(tc.grid, x, y, visited)
			if !reflect.DeepEqual(visited, tc.want) {
				t.Errorf("got: %v, want: %v", visited, tc.want)
			}
		})
	}
}

func TestGetAdjacent(t *testing.T) {
	tests := map[string]struct {
		grid  [][]int
		input []int
		want  [][]int
	}{
		"1": {
			grid: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			input: []int{0, 0}, // 座標0, 0の隣接点は、
			want: [][]int{
				{1, 0}, // 座標1, 0
				{0, 1}, // 座標0, 1
			},
		},
		"2": {
			grid: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			input: []int{1, 1}, // 座標1, 1の隣接点は、
			want: [][]int{
				{0, 1},
				{1, 2},
				{1, 0},
			},
		},
		"3": {
			grid: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 1},
				{0, 0, 0, 1, 1},
			},
			input: []int{4, 3},
			want: [][]int{
				{3, 3},
				{4, 2},
			},
		},
		"4": {
			grid: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			input: []int{3, 3},
			want: [][]int{
				{4, 3},
				{3, 2},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			x, y := tc.input[0], tc.input[1]
			got := getAdjacent(tc.grid, x, y)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

type Visited [][]int
type Visiteds []Visited

func maxAreaOfIsland(grid [][]int) int {
	var max int

	var visiteds Visiteds
	for y, v := range grid {
		for x, vv := range v {
			if vv == 1 {
				if isInVisiteds(x, y, visiteds) {
					continue
				}

				var visited Visited
				visited = dfs(grid, x, y, visited)
				// fmt.Println("x, y", x, y, len(visited))
				if max < len(visited) {
					max = len(visited)
				}

				visiteds = append(visiteds, visited)
			}
		}
	}

	return max
}

func isInVisiteds(x, y int, visiteds Visiteds) bool {
	for _, visited := range visiteds {
		if isVisited(x, y, visited) {
			return true
		}
	}
	return false
}

func dfs(grid [][]int, x, y int, visited [][]int) [][]int {
	if grid[y][x] == 0 {
		return visited
	}

	visited = append(visited, []int{x, y})
	for _, xy := range getAdjacent(grid, x, y) {
		x, y = xy[0], xy[1]

		if isVisited(x, y, visited) {
			continue
		}
		visited = dfs(grid, x, y, visited)
	}

	return visited
}

func isVisited(x, y int, visited [][]int) bool {
	for _, xy := range visited {
		if xy[0] == x && xy[1] == y {
			return true
		}
	}
	return false
}

// 対象の座標を与えたら、隣接してかつ１である点の座標を返す関数
func getAdjacent(grid [][]int, i, j int) [][]int {

	candidate := [][]int{
		{i + 1, j},
		{i - 1, j},
		{i, j + 1},
		{i, j - 1},
	}

	var filtered [][]int
	for _, v := range candidate {
		if isValidGrid(v, len(grid[0]), len(grid)) {
			filtered = append(filtered, v)
		}
	}

	var adjacent [][]int
	for _, v := range filtered {

		if grid[v[1]][v[0]] == 1 {
			adjacent = append(adjacent, v)
		}
	}

	return adjacent
}

// 座標が有効だったらtrue
func isValidGrid(b []int, xMax, yMax int) bool {
	if b[0] < 0 || b[1] < 0 || b[0] >= xMax || b[1] >= yMax {
		return false
	}
	return true
}
