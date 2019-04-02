# MapToVec

![](https://goreportcard.com/badge/github.com/vincentfiestada/maptovec)

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
```
