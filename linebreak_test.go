package linebreak_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/invisiblefunnel/linebreak"
)

func planarDist(a, b [2]float64) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(dx*dx + dy*dy)
}

func TestSolveBreaks(t *testing.T) {
	var targetDist float64 = 11

	points := [][2]float64{{0, 0}, {3, 8}, {4, 16}, {10, 23}, {11, 28}}

	dists := make([]float64, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		dists[i] = planarDist(points[i], points[i+1])
	}

	var actual [][]int
	linebreak.Solve(dists, targetDist, func(i, j int) {
		actual = append(actual, []int{i, j})
	})

	expected := [][]int{{0, 1}, {1, 2}, {2, 4}}

	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}

func TestSolveBreaksNone(t *testing.T) {
	points := [][2]float64{{0, 0}, {3, 8}, {4, 16}, {10, 23}, {11, 28}}

	var totalDistance float64
	dists := make([]float64, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		dist := planarDist(points[i], points[i+1])
		dists[i] = dist
		totalDistance += dist
	}

	var actual [][]int
	linebreak.Solve(dists, totalDistance, func(i, j int) {
		actual = append(actual, []int{i, j})
	})

	expected := [][]int{{0, len(points) - 1}}

	if len(expected) != 1 {
		t.Fail()
	}

	if expected[0][0] != 0 {
		t.Fail()
	}

	if expected[0][1] != len(points)-1 {
		t.Fail()
	}
}

func TestSolveBreaksAll(t *testing.T) {
	points := [][2]float64{{0, 0}, {3, 8}, {4, 16}, {10, 23}, {11, 28}}

	minDistance := math.Inf(1)
	dists := make([]float64, len(points)-1)
	for i := 0; i < len(points)-1; i++ {
		dist := planarDist(points[i], points[i+1])
		dists[i] = dist
		if dist < minDistance {
			minDistance = dist
		}
	}

	var actual [][]int
	linebreak.Solve(dists, minDistance, func(i, j int) {
		actual = append(actual, []int{i, j})
	})

	var expected [][]int
	for i := 0; i < len(points)-1; i++ {
		expected = append(expected, []int{i, i + 1})
	}

	if len(expected) != len(points)-1 {
		t.Fail()
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fail()
	}
}
