package Integers

func GetNthFib(n int) int {
	// Write your code here.
	previous, current := 0, 1
	for i := 0; i < n-1; i++ {
		previous, current = current, previous+current
	}
	return previous
}
