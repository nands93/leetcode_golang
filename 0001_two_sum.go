package main

import "fmt"

func twoSum(nums []int, target int) []int {
    var output []int
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if (nums[i] + nums[j] == target) {
                output = append(output, i, j)
                return output
            }
        } 
    }
    return nil
}

func main() {
	test1 := []int{2, 7, 11, 15}
	test2 := []int{3, 2, 4}
	test3 := []int{3, 3}
	test4 := []int{2, 5, 5, 11}
	target1 := 9
	target2 := 6
	target3 := 6
	target4 := 10
	fmt.Println("Teste 1: ", twoSum(test1, target1))
	fmt.Println("Teste 2: ", twoSum(test2, target2))
	fmt.Println("Teste 3: ", twoSum(test3, target3))
	fmt.Println("Teste 4: ", twoSum(test4, target4))
}
