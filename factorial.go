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
    num, err := strconv.Atoi(os.Args[1]) // converts string to int

    // Functions return your return and an error. If it's nil == error!
    if err != nil {
        fmt.Println("Error! ", err)
    } else {
        var result int = factorial(num)
        fmt.Println("Factorial of", num, ":", result)
    }
}
