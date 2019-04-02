# MapToVec

[![Build Status](https://vincentofearth.visualstudio.com/maptovec/_apis/build/status/maptovec%20Continuous%20Integration?branchName=master)](https://vincentofearth.visualstudio.com/maptovec/_build/latest?definitionId=4&branchName=master) [![](https://goreportcard.com/badge/github.com/vincentfiestada/maptovec)](https://goreportcard.com/report/github.com/vincentfiestada/maptovec) [![GoDoc](https://godoc.org/github.com/vincentfiestada/maptovec?status.svg)](https://godoc.org/github.com/vincentfiestada/maptovec)

Convert a Map with string keys to a float64 array. Numeric, boolean, and nil values are converted directly to float64. For other types, each new value encountered is given an index unique for its key.

A **Vectorizer** can be used to convert maps with a uniform set of keys, whose values are all [comparable](https://golang.org/ref/spec#Comparison_operators).

```go
v := maptovec.NewVectorizer()

v.MapToVec(maptovec.Map{
	"a": 25,
	"b": float32(2.5),
	"c": float64(2.5),
	"d": true,
	"e": false,
	"f": "25",
	"g": nil,
}) // returns [25 2.5 2.5 1 0 1 0]

v.MapToVec(maptovec.Map{
	"a": 75,
	"b": float32(7.5),
	"c": float64(7.5),
	"d": true,
	"e": false,
	"f": "26",
	"g": struct{a string}{
		a: "abc",
	},
}) // returns [75 7.5 7.5 1 0 2 2]
```
