package main

import (
    "fmt"
    "errors"
)

// Simple recursion
func fact(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * fact(n - 1)
    }
}

// Error handling
func maybe_error(n int) (int, error) {
    if n % 2 == 0 {
        return n/2, nil
    }
    return 0, errors.New("Odd number; expected even")
}

func mysum_int(numbers ...int) int {
    ans := 0
    for _, val := range numbers {
        ans += val
    }
    return ans
}

func defer_test() {
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")
    defer fmt.Println("Defer 3")
    fmt.Println("Defer 4")
}

// 5-element int array arg and return value
func map_int1(lst [5]int, f func(int) int) [5]int {
    var result_list [5]int
    for index, item := range(lst) {
        result_list[index] = f(item)
    }
    return result_list
}

// Slice arg and return value
func map_int2(lst []int, f func(int) int) []int {
    result_list := make([]int, 0, len(lst))
    for _, item := range(lst) {
        result_list = append(result_list, f(item))
    }
    return result_list
}

// Generic slice arg and return value - not very nice to use in Go
func map_func(lst []interface{}, f func(interface{}) interface{}) []interface{} {
    result_list := make([]interface{}, 0, len(lst))
    for _, item := range(lst) {
        result_list = append(result_list, f(item))
    }
    return result_list
}

func main() {
    fmt.Println("======= Simple function =======")
    fmt.Println(fact(5))

    fmt.Println("======= Error handling =======")
    if doubled, err := maybe_error(3); err != nil {
        fmt.Println("(Expeted) Error:", err)
    } else {
        fmt.Println(doubled)
    }

    fmt.Println("======= Variable length arguments =======")
    fmt.Println("Sum 1..5:", mysum_int(1, 2, 3, 4, 5))

    fmt.Println("======= Defer =======")
    defer_test()

    fmt.Println("======= Map test 1 =======")
    fmt.Println(map_int1([...]int{1, 2, 3, 4, 5}, func (num int) int { return num * num }))

    fmt.Println("======= Map test 2 =======")
    fmt.Println(map_int2([]int{1, 2, 3, 4, 5, 6, 7}, func (num int) int { return num * num }))

    fmt.Println("======= Map test 3 =======")
    // Ick!
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    numbers_interface := make([]interface{}, len(numbers))
    for i, v := range numbers {
        numbers_interface[i] = v
    }
    fmt.Println(map_func(numbers_interface, func (val interface{}) interface{} { return val.(int) * val.(int) }))
}
