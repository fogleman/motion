package motion

import "math"

type triangle struct {
	s1, s2     float64
	t1, t2     float64
	vmax       float64
	p1, p2, p3 Point
}

// triangularProfile computes a triangular profile: accelerating, decelerating
func triangularProfile(s, vi, vf, a float64, p1, p3 Point) triangle {
	s1 := (2*a*s + vf*vf - vi*vi) / (4 * a)
	s2 := s - s1
	vmax := math.Sqrt(vi*vi + 2*a*s1)
	t1 := (vmax - vi) / a
	t2 := (vf - vmax) / -a
	p2 := p1.lerps(p3, s1)
	return triangle{s1, s2, t1, t2, vmax, p1, p2, p3}
}
