package main

import (
	"fmt"
	"math"
)

func main() {
	var W, H, x, y float64
	fmt.Scan(&W, &H, &x, &y)

	area := W * H
	xSplits := []float64{x / W, 1 - (x / W)}
	ySplits := []float64{y / H, 1 - (y / H)}
	fmt.Println(area)
	xDiff := math.Abs(xSplits[0] - xSplits[1])
	yDiff := math.Abs(ySplits[0] - ySplits[1])

	var res float64 = 0
	if xDiff >= yDiff {
		if ySplits[0] >= ySplits[1] {
			res = area * ySplits[1]
		} else {
			res = area * ySplits[0]
		}
	}

	if xDiff < yDiff {
		if xSplits[0] >= xSplits[1] {
			res = area * xSplits[1]
		} else {
			res = area * xSplits[0]
		}
	}

	count := 0
	if x == y {
		count = 1
	}
	fmt.Printf("%f %d\n", res, count)

}
