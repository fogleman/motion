package motion

import (
	"fmt"
	"testing"
)

func TestPlan(tt *testing.T) {
	var points []Point
	points = append(points, Point{0, 0})
	points = append(points, Point{100, 0})
	points = append(points, Point{0, 0})

	var velocities []float64
	velocities = append(velocities, 0)
	velocities = append(velocities, 100)
	velocities = append(velocities, 0)

	plan := NewPlan(points, velocities, 1, 10, 0.1)

	var t float64
	for t < plan.T {
		fmt.Println(plan.Instant(t))
		t += 1
	}
}
