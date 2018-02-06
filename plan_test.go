package motion

import (
	"fmt"
	"testing"
)

func TestPlan(tt *testing.T) {
	var points []Point
	points = append(points, Point{0, 0})
	points = append(points, Point{1, 0})
	points = append(points, Point{2, 0})
	points = append(points, Point{3, 0})
	points = append(points, Point{4, 0})

	var vs []float64
	vs = append(vs, 0)
	vs = append(vs, 10)
	vs = append(vs, 10)
	vs = append(vs, 10)
	vs = append(vs, 0)

	var vmaxs []float64
	vmaxs = append(vmaxs, 5)
	vmaxs = append(vmaxs, 10)
	vmaxs = append(vmaxs, 5)
	vmaxs = append(vmaxs, 1)
	vmaxs = append(vmaxs, 10)

	plan := NewPlan(points, vs, vmaxs, 20, 10, 0.1)

	var t float64
	for t < plan.T {
		x := plan.Instant(t)
		fmt.Printf("%6.3f %6.3f %6.3f %6.3f\n", x.T, x.S, x.V, x.A)
		t += 0.01
	}
}
