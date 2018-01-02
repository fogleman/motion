package motion

type trapezoid struct {
	s1, s2, s3     float64
	t1, t2, t3     float64
	p1, p2, p3, p4 Point
}

// trapezoidalProfile computes a trapezoidal profile: accelerating, cruising,
// decelerating
func trapezoidalProfile(s, vi, vmax, vf, a float64, p1, p4 Point) trapezoid {
	t1 := (vmax - vi) / a
	s1 := (vmax + vi) / 2 * t1
	t3 := (vf - vmax) / -a
	s3 := (vf + vmax) / 2 * t3
	s2 := s - s1 - s3
	t2 := s2 / vmax
	p2 := p1.lerps(p4, s1)
	p3 := p1.lerps(p4, s-s3)
	return trapezoid{s1, s2, s3, t1, t2, t3, p1, p2, p3, p4}
}
