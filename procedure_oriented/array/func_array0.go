package main

import "fmt"

func main() {

	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32


	avg = getAverage(balance, 5)


	fmt.Printf("平均值为: %f \n", avg)

	for i := 0; i < 5; i++ {
		fmt.Println(balance[i]);
	}
}

// copy
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		arr[i] = 1;
	}

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
