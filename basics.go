package main

import "os"
import "fmt"
import "time"
import "math/rand"

func main() {
    // Command-line args
    name := "friend"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }

    // Output
    fmt.Println("Hello", name)

    // Time
    t := time.Now()
    fmt.Println("Time is", t)

    // Random
    rand.Seed(t.UnixNano() / int64(time.Millisecond))
    fmt.Println("Random:", rand.Intn(10))

    // Variables
    var s string = "Caleb Madrigal"
    s2 := "Caleb Madrigal"  // Shorthand for declare and initialize
    fmt.Println(s == s2)

    // Loops
    i := 0
    for i < 10 {
        i = i + 1
    }

    for j := 0; j < 10; j++ {
        i += 1
    }
    fmt.Println("i =", i)

    // Arrays
    numbers := [...]int{1, 2, 3, 4, 5}
    numbers[1] = 20
    fmt.Println(numbers[1:4])

    numbers2 := append(numbers[:], 6, 7, 8, 9)

    // Foreach
    for _, value := range numbers2 {
        fmt.Println("val =", value)
    }

    // Maps
    var name_age = make(map[string]int)
    name_age["caleb"] = 53225
    name_age["whitney"] = 53225
    name_age["dan"] = 62278
    name_age["test"] = 12345
    delete(name_age, "test")
    fmt.Println(name_age)

}

