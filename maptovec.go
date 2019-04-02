package maptovec

import "sort"

// Map is a string to (any) mapping
type Map map[string]interface{}

// Vec is a vector of real numbers
type Vec []float64

// index is a mapping from (any) to float
type index map[interface{}]float64

// Vectorizer is a utility for converting Maps to Vecs
type Vectorizer struct {
	indexes map[string]index
	nextIDs map[string]float64
}

// NewVectorizer creates and returns a new Vectorizer
func NewVectorizer() *Vectorizer {
	vectorizer := &Vectorizer{}
	vectorizer.indexes = make(map[string]index)
	vectorizer.nextIDs = make(map[string]float64)
	return vectorizer
}

// nextID returns the next available id for a given index name
func (vectorizer *Vectorizer) nextID(name string) float64 {
	vectorizer.nextIDs[name]++
	return vectorizer.nextIDs[name]
}

// indexMap adds a map's key-value pairs to the appropriate indices
func (vectorizer *Vectorizer) indexMap(src Map) {
	for _, key := range src.orderedKeys() {
		elem := src[key]
		if vectorizer.indexes[key] == nil {
			vectorizer.indexes[key] = index{}
		}
		if vectorizer.indexes[key][elem] == 0 {
			vectorizer.indexes[key][elem] = vectorizer.nextID(key)
		}
	}
}

// MapToVec converts a Map to a Vec
func (vectorizer *Vectorizer) MapToVec(src Map) Vec {
	vectorizer.indexMap(src)
	vec := Vec{}
	for _, key := range src.orderedKeys() {
		elem := src[key]
		switch elem.(type) {
		case float64:
			vec = append(vec, elem.(float64))
		case float32:
			vec = append(vec, float64(elem.(float32)))
		case int:
			vec = append(vec, float64(elem.(int)))
		case bool:
			vec = append(vec, boolToFloat(elem.(bool)))
		case nil:
			vec = append(vec, 0)
		default:
			vec = append(vec, vectorizer.indexes[key][elem])
		}
	}
	return vec
}

// orderedKeys returns the Map's keys, in alphabetical order
func (m Map) orderedKeys() []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
