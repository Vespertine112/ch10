// Example package which exposes an addition function
package ch10

type Number interface{
	~int | ~float32 | ~float64
}

// Add
// Sums two ints together and returns an int
//
// Refer to [Addition] for more information
//
// [Addition]: https://mathisfun.com/numbers/addition.html
func Add[T Number](a,b T) T {
	 return a + b
}

