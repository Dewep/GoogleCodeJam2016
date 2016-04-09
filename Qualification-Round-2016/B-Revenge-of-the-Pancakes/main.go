package main

import "fmt"

func pancakes(S string) int {
    length := 0
    last_minus := 0

    for i := 0; i < len(S); i++ {
        if i == 0 || S[i] != S[i - 1] { // Group operators (---++--+ => -+-+)
            length += 1

            if S[i] == '-' { // Last position of minus is the number of flips
                last_minus = length
            }
        }
    }

    return last_minus
}

func main() {
    var T int
    fmt.Scan(&T)

    for i := 1; i <= T; i++ {
        var S string
        fmt.Scan(&S)

        times := pancakes(S)

        fmt.Printf("Case #%d: %d\n", i, times)
    }
}
