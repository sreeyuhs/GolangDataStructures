// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//O(n) - Time Complexity
//O(1) - Space Cpmplexity

func MinMax(A []int) (int, int) {
	min, max := A[0], A[0]
	for i := 0; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	return min, max
}

func main() {
	A := []int{1, 2, 4, 12, 22, 84, 111, 1232}
	fmt.Println(MinMax(A))
}
