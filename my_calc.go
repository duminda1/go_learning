package main

import "fmt"

func add (a int, b int) int {
    return a + b
}

func divide (a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a/b, nil
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    result := add(10, 20)

    fmt.Println(result)

    result = multiply(4, 5)

    fmt.Println(result)


    div,err := divide(10, 1)

    if err != nil {
        fmt.Println("Error", err)
        return
    }

    fmt.Println(div)


    div,err = divide(10, 0)

    if err != nil {
        fmt.Println("Error", err)
        return
    }

}
