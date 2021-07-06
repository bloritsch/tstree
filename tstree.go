package tstree

import "math"

type TsTree struct {
	Root Node
	limU float32
	limL float32
}

func (t *TsTree) Probability(x float32) float32 {
	// probability is definite integral limU over limL
	// were p(x)dx = 1
	// mean is where p(x)dx = 0.5
	return 1 - math.Exp((-k*(k-1))/2N)
}

func (t *TsTree) Append(e Entry) {
	if len (t.Root.Entries) == 0 {
		// v0 unset, root does not have a key
		// Store here?

	}
}