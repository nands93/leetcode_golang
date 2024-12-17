package main

import "fmt"

func isPalindrome(x int) bool {
    palindrome := 0
    original := x
    for x > 0 {
        palindrome = palindrome * 10 + (x % 10)
        x /= 10
    }
    if original == palindrome {
        return true
    }
    return false
}

func main() {
	test1 := 121
	test2 := -121
	test3 := 10
	test4 := 0
	fmt.Println("Teste 1: ", isPalindrome(test1))
	fmt.Println("Teste 2: ", isPalindrome(test2))
	fmt.Println("Teste 3: ", isPalindrome(test3))
	fmt.Println("Teste 4: ", isPalindrome(test4))
}