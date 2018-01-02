package motion

import (
	"math"
	"sort"
)

type Plan struct {
	Blocks []*Block
	T, S   float64
	Ts, Ss []float64
}

func NewPlan(points []Point, velocities []float64, a, vmax, cf float64) *Plan {
	const eps = 1e-9

	// create segments for each consecutive pair of points
	segments := make([]*segment, 0, len(points))
	for i := 1; i < len(points); i++ {
		segments = append(segments, newSegment(points[i-1], points[i]))
	}

	// add a dummy segment at the end for setting final velocity
	lastPoint := points[len(points)-1]
	segments = append(segments, newSegment(lastPoint, lastPoint))

	if len(velocities) == 0 {
		// compute a maxEntryVelocity for each segment based on the angle
		// formed by the two segments and the cornering factor (cf)
		// the first and last segments will have maxEntryVelocity == 0
		for i := 1; i < len(segments)-1; i++ {
			s1 := segments[i-1]
			s2 := segments[i]
			s2.maxEntryVelocity = cornerVelocity(s1, s2, vmax, a, cf)
		}
	} else if len(velocities) == len(points) {
		// set maxEntryVelocity based on input
		for i, v := range velocities {
			segments[i].maxEntryVelocity = math.Min(v, vmax)
		}
	} else {
		panic("velocities array must be empty or same length as points array")
	}

	// loop over segments
	i := 0
	for i < len(segments)-1 {
		// pull out some variables
		segment := segments[i]
		nextSegment := segments[i+1]
		s := segment.length
		vi := segment.entryVelocity
		vexit := nextSegment.maxEntryVelocity
		p1 := segment.p1
		p2 := segment.p2
		blocks := segment.blocks[:0]

		// determine which profile to use for this segment
		m := triangularProfile(s, vi, vexit, a, p1, p2)
		if m.s1 < -eps {
			// too fast! update max_entry_velocity and backtrack
			segment.maxEntryVelocity = math.Sqrt(vexit*vexit + 2*a*s)
			i--
		} else if m.s2 < 0 {
			// accelerate
			vf := math.Sqrt(vi*vi + 2*a*s)
			t := (vf - vi) / a
			blocks = append(blocks, NewBlock(a, t, vi, p1, p2))
			nextSegment.entryVelocity = vf
			i++
		} else if m.vmax > vmax {
			// accelerate, cruise, decelerate
			z := trapezoidalProfile(s, vi, vmax, vexit, a, p1, p2)
			blocks = append(blocks, NewBlock(a, z.t1, vi, z.p1, z.p2))
			blocks = append(blocks, NewBlock(0, z.t2, vmax, z.p2, z.p3))
			blocks = append(blocks, NewBlock(-a, z.t3, vmax, z.p3, z.p4))
			nextSegment.entryVelocity = vexit
			i++
		} else {
			// accelerate, decelerate
			blocks = append(blocks, NewBlock(a, m.t1, vi, m.p1, m.p2))
			blocks = append(blocks, NewBlock(-a, m.t2, m.vmax, m.p2, m.p3))
			nextSegment.entryVelocity = vexit
			i++
		}
		segment.blocks = blocks
	}

	// concatenate all of the blocks
	var blocks []*Block
	for _, segment := range segments {
		for _, block := range segment.blocks {
			if block.T > eps {
				blocks = append(blocks, block)
			}
		}
	}

	// compute starting time and position for each block
	ts := make([]float64, len(blocks))
	ss := make([]float64, len(blocks))
	var t, s float64
	for i, block := range blocks {
		ts[i] = t
		ss[i] = s
		t += block.T
		s += block.S
	}

	return &Plan{blocks, t, s, ts, ss}
}

func (p *Plan) Instant(t float64) Instant {
	t = clamp(t, 0, p.T)
	i := sort.Search(len(p.Ts), func(i int) bool { return p.Ts[i] > t }) - 1
	return p.Blocks[i].Instant(t-p.Ts[i], p.Ts[i], p.Ss[i])
}

func (p *Plan) InstantAtDistance(s float64) Instant {
	s = clamp(s, 0, p.S)
	i := sort.Search(len(p.Ss), func(i int) bool { return p.Ss[i] > s }) - 1
	return p.Blocks[i].InstantAtDistance(s-p.Ss[i], p.Ts[i], p.Ss[i])
}
