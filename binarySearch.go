// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//Binary Search - To get a Peak Point/Midpoint which such that elements to its left are smaller and elements to its right are bigger, if they exist.
//Use Recursion - O(LogN) - time and space complexity

func BinarySearch(A []int, low int, high int) int {
	mid := low + (high-low)/2
	n := len(A)
	fmt.Println("mid is ", mid)

	if (mid == 0 || A[mid-1] <= A[mid]) && (mid == n-1 || A[mid] < A[mid+1]) {
		return mid
	} else if A[mid-1] <= A[mid] {
		return BinarySearch(A, low, mid-1)
	} else {
		return BinarySearch(A, mid+1, n-1)
	}
}

func main() {
	A := []int{1, 2, 4, 12, 22, 84, 111, 1232}
	fmt.Println(BinarySearch(A, 0, len(A)))
}
