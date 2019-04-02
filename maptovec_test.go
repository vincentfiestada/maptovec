package maptovec

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestMapToVec(t *testing.T) {
	g := Goblin(t)
	g.Describe("MapToVec()", func() {
		g.It("Should return a Vec with same len as Map", func() {
			v := NewVectorizer()
			input := makeInput()
			g.Assert(len(v.MapToVec(input))).Equal(len(input))
		})
		g.It("Should return different Vecs for different Maps", func() {
			v := NewVectorizer()
			a := v.MapToVec(Map{
				"a": "abc",
			})
			b := v.MapToVec(Map{
				"a": "xyz",
			})
			for i := range a {
				g.Assert(a[i] == b[i]).IsFalse()
			}
		})
		g.It("Should be deterministic", func() {
			v := NewVectorizer()
			a := v.MapToVec(makeInput())
			b := v.MapToVec(makeInput())
			for i := range a {
				g.Assert(a[i]).Equal(b[i])
			}
		})
	})
}

// makeInput returns a new Map
func makeInput() Map {
	return Map{
		"a": 25,
		"b": float32(2.5),
		"c": float64(2.5),
		"d": true,
		"e": false,
		"f": "25",
		"g": nil,
	}
}
