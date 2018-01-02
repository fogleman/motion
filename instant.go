package motion

type Instant struct {
	T float64 // time elapsed
	S float64 // distance traveled
	V float64 // instantaneous velocity
	A float64 // instantaneous acceleration
	P Point   // instantaneous position
}
