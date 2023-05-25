package linebreak

import "math"

type solver struct {
	distMatrix [][]float64
	targetDist float64

	n              int
	memo           []float64
	parentPointers []int
}

func Solve(dists []float64, targetDist float64, f func(i, j int)) {
	n := len(dists) + 1
	memo := make([]float64, n)
	distMatrix := make([][]float64, n)

	for i := 0; i < n; i++ {
		memo[i] = -1
		distMatrix[i] = make([]float64, n)
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			distMatrix[i][i+j] = distMatrix[i][i+j-1] + dists[i+j-1]
		}
	}

	s := solver{
		distMatrix: distMatrix,
		targetDist: targetDist,

		n:              n,
		memo:           memo,
		parentPointers: make([]int, n),
	}

	s.solve(f)
}

func (s solver) solve(f func(i, j int)) {
	s.dp(0)

	var start, end int
	for start != s.n-1 {
		end = s.parentPointers[start]
		f(start, end)
		start = end
	}
}

func (s solver) dp(i int) float64 {
	if s.memo[i] >= 0 {
		return s.memo[i]
	}

	if i == s.n-1 {
		s.memo[i] = 0
		return s.memo[i]
	}

	minBadness := math.Inf(1)
	for j := i + 1; j < s.n; j++ {
		badness := s.badness(i, j) + s.dp(j)
		if badness < minBadness {
			minBadness = badness
			s.parentPointers[i] = j
		}
	}

	s.memo[i] = minBadness
	return minBadness
}

func (s solver) badness(i, j int) float64 {
	return math.Pow(s.distMatrix[i][j]-s.targetDist, 2)
}
