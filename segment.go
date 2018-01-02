package motion

type segment struct {
	p1, p2           Point
	vector           Point
	length           float64
	maxEntryVelocity float64
	entryVelocity    float64
	blocks           []*Block
}

func newSegment(p1, p2 Point) *segment {
	vector := p2.sub(p1).normalize()
	length := p1.distance(p2)
	return &segment{p1, p2, vector, length, 0, 0, nil}
}
