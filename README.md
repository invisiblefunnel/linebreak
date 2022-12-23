# linebreak

Split a sequence of points so that each subsequence is close to a desired length. This is useful, e.g., if you want to break a geographic LineString into a MultiLineString _without_ interpolating new points or altering existing points. The algorithm is an adaptation of text justification using dynamic programming.

## Usage

```golang
package main

import (
    "fmt"
    "math"

    "github.com/invisiblefunnel/linebreak"
)

func main() {
    line := [][2]float64{{0, 0}, {3, 8}, {4, 16}, {10, 23}, {11, 28}}

    dists := make([]float64, len(line)-1)
    for i := 0; i < len(line)-1; i++ {
        dists = append(dists, pointDist(line[i], line[i+1]))
    }

    targetDist := 11.0
    linebreak.Solve(dists, targetDist, func(i, j int) {
        fmt.Println(lineDist(line[i:j+1]), line[i:j+1])
    })
}

func pointDist(a, b [2]float64) float64 {
    dx := a[0] - b[0]
    dy := a[1] - b[1]
    return math.Sqrt(dx*dx + dy*dy)
}

func lineDist(points [][2]float64) float64 {
    var length float64
    for i := 0; i < len(points)-1; i++ {
        length += pointDist(points[i], points[i+1])
    }
    return length
}
```
