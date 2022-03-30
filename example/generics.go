package example

/*
Proof of concept for a golang batch processor using
generics and a CRUD api.

*/
// func main() {

// 	// Initialize a map for the integer values
// 	ints := map[string]int64{
// 		"first":  34,
// 		"second": 12,
// 	}

// 	// Initialize a map for the float values
// 	floats := map[string]float64{
// 		"first":  35.98,
// 		"second": 26.99,
// 	}
// 	fmt.Printf("Non-Generic Sums: %v and %v\n",
// 		SumInts(ints),
// 		SumFloats(floats))

// 	// in what case would you NEED to specify the types?
// 	fmt.Printf("Generic Sums: %v of type %T and %v of type %T\n",
// 		SumIntsOrFloats[string, int64](ints),
// 		SumIntsOrFloats[string, int64](ints),
// 		SumIntsOrFloats[string, float64](floats),
// 		SumIntsOrFloats[string, float64](floats))

// 	fmt.Printf("Generic Sums, type parameters inferred: %v of type %T and %v of type %T\n",
// 		SumIntsOrFloats(ints),
// 		SumIntsOrFloats(ints),
// 		SumIntsOrFloats(floats),
// 		SumIntsOrFloats(floats))
// }

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
