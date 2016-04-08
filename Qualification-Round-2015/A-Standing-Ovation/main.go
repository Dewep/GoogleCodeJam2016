package main

import "fmt"

func solve(Smax int, levels string) string {
    fmt.Println(Smax)
    fmt.Println(levels)
    fmt.Println("---")
    return "toto"
}

func main() {
    var T int
    fmt.Scan(&T)
    for i := 1; i <= T; i++ {
        var Smax int
        var levels string
        fmt.Scan(&Smax)
        fmt.Scan(&levels)
        answer := solve(Smax, levels)
        fmt.Printf("Case #%d: %d\n", i, answer)
    }
}
