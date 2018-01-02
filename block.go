package motion

type Block struct {
	A  float64
	T  float64
	Vi float64
	S  float64
	P1 Point
	P2 Point
}

func NewBlock(a, t, vi float64, p1, p2 Point) *Block {
	s := p1.distance(p2)
	return &Block{a, t, vi, s, p1, p2}
}

func (b *Block) Instant(t, dt, ds float64) Instant {
	t = clamp(t, 0, b.T)
	a := b.A
	v := b.Vi + b.A*t
	s := b.Vi*t + b.A*t*t/2
	s = clamp(s, 0, b.S)
	p := b.P1.lerps(b.P2, s)
	return Instant{t + dt, s + ds, v, a, p}
}
