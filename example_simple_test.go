package got_test

import (
	"testing"

	"github.com/ysmood/got"
)

// The simplest form to use got.Each
func TestSimple(t *testing.T) {
	got.Each(t, Simple{})
}

type Simple struct {
	got.Assertion
	*testing.T
}

func (s Simple) A() {
	s.Eq(1, 1)
}

func (s Simple) B() {
	s.Skip()

	s.Gt(2, 1)
}
