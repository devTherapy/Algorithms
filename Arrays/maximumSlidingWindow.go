package arrays

// MaximumSlidingWindow returns the maximum value in a sliding window of size k.
//
// Given an array of integers and an integer k, return the maximum value in each sliding window of size k.
//
// Parameters:
//   array []int - The array of integers.
//   windowSize int - The size of the sliding window.
//
// Returns:
//   []int - An array of integers representing the maximum value in each sliding window of size k.
//
//TC O(N), SC O(N)
//
//This function works by using a queue to store the indexes of the elements in the array. We iterate through the array and add the index of the element to the queue if the element is greater than the last element in the queue. We then check if the index of the first element in the queue is less than the beginning of the new window and remove it from the queue. If the index of the current element is greater than or equal to the window size, we add the maximum value in the window to the result array.
//
//The logic here is that we only need to keep track of the maximum value in the window and remove any elements that are less than the current element since they will never be the maximum value in the next window.
func MaximumSlidingWindow(array []int, windowSize int) []int {
	queue := []int{}
	result := []int{}
	//we use l to track the beginning of the new window
	l:= 0
	for i, v :=  range array{
		//while current value is greater than what we adding into the queue
		for len(queue) > 0 && array[queue[len(queue) - 1]] < v{
			queue = queue[:len(queue) -1]  //remove the last element from the slice
		}
		queue = append(queue, i)
		//if the index of the first element in the queue is less than the beginning of the new window, remove it from the queue
		if l > queue[0] {
			queue = queue[1:]
		}
		//if we have reached a valid window size, add the maximum value in the window to the result array
		indexInValidWindow := i + 1 >= windowSize
		if indexInValidWindow {
			result = append(result, array[queue[0]])
			l++
		}

	}

	return result
}