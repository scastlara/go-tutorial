package main

import "fmt"
import "os"
import "strconv"

// (args) (returns)
func factorial(n int)(int) {
    if n == 0 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}

func main() {
    num, _ := strconv.Atoi(os.Args[1]) // converts string to int
    var result int = factorial(num)
    fmt.Println("Factorial of", num, ":", result)
}
