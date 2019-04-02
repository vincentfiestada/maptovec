package maptovec

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestConvert(t *testing.T) {
	g := Goblin(t)
	g.Describe("boolToFloat", func() {
		g.It("Should return 1.0 for true", func() {
			g.Assert(boolToFloat(true)).Equal(1.0)
		})
		g.It("Should return 0.0 for false", func() {
			g.Assert(boolToFloat(false)).Equal(0.0)
		})
	})
}
