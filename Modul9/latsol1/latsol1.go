package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    
    k := 2
    m := n / k

    
    if n % k != 0 {
        m++
    }

    fmt.Print(m)
}
