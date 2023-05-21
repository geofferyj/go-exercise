package main

import "fmt"

func main() {

	arr:= []int{1,2,3,2,1,2,3,4,3,4,5,6,5,5,677,65,44,3,332,2}
	selectionSort(arr)

}

func selectionSort(arr []int) []int{
	fmt.Println("before sort", arr)

	for i:=0; i < len(arr); i++{
		minIndex:= i

		for index, value:= range arr[i:]{
			if arr[minIndex] > value {
				minIndex = index + i
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}