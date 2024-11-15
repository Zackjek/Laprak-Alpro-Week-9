package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    xFaktorY := false
    if y % x == 0 {
        xFaktorY = true
    }

    yFaktorX := false
    if x % y == 0 {
        yFaktorX = true
    }

    fmt.Println(xFaktorY)
    fmt.Println(yFaktorX)
}
