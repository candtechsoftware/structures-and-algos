package main

import "fmt"

func isValiSubsequence(arr []int, seq []int) bool {
	index := 0

	/*
		Looping Through the array (Not the Subsequence)`
		I will use the index varaible to see where I am in the seqeuence array
		if the value in the array is equal to the value in the index location
		of the seqeuence I will move the index up. With that logic for a valid sequence
		the index should equal the lenghth of the sequence array. If it happens during the loop I break out since
		there is no need to continue iterating over the larger array.
	*/
	for _, val := range arr {

		if index == len(seq) {
			break
		}
		if val == seq[index] {
			index++
		}
	}

	return index == len(seq)
}

func main() {
	ans1, ans2 := isValiSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}), isValiSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 3, 33, 13})

	fmt.Println(ans1, ans2)
}
