package motion

import "math"

// cornerVelocity computes a maximum velocity at the corner of two segments
// https://onehossshay.wordpress.com/2011/09/24/improving_grbl_cornering_algorithm/
func cornerVelocity(s1, s2 *segment, vmax, a, cf float64) float64 {
	const eps = 1e-9
	cosine := -s1.vector.dot(s2.vector)
	if math.Abs(cosine-1) < eps {
		return 0
	}
	sine := math.Sqrt((1 - cosine) / 2)
	if math.Abs(sine-1) < eps {
		return vmax
	}
	v := math.Sqrt((a * cf * sine) / (1 - sine))
	return math.Min(v, vmax)
}

func clamp(x, x0, x1 float64) float64 {
	if x < x0 {
		x = x0
	}
	if x > x1 {
		x = x1
	}
	return x
}
