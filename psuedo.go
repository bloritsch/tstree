package tstree

// Range l, u
const (
	// 26 characters
	l         = "00000000000000000000000000"
	u         = "7ZZZZZZZZZZZZZZZZZZZZZZZZZ"
	increment = 1e-5
)

func FindMLeft(r Node, v string) Entry {
	if r.Right != nil {
		return FindMLeft(*r.Right, v)
	} else if e := r.Find(v); e != nil {
		return *e
	} else if r.Left != nil {
		return FindMLeft(*r.Left, v)
	}

	return Entry{
		Key: l,
	}
}

func Distribution(u, l) float32 {
	// definite integral between u, l
	// p(x)dx = 1
}

// ProbabilityLeft() + ProbabilityRight() = 1
// ProbabilityDiff = abs(ProbabilityLeft() - ProbabilityRight()) = abs(2*ProbabilityLeft() - 1)

func integral(a float32, b float32, f func(x float32) float32) float32 {
	area := float32(0.0)
	modifier := float32(1.0)

	if a > b {
		temp := a
		a = b
		b = temp
		modifier = -1
	}

	for i := a + increment; i < b; i += increment {
		distanceA := i - a
		area += (increment/2)*f(a+distanceA) + f(a+distanceA-increment)
	}

	return area * modifier
}
