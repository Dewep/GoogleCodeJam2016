package main

import "fmt"
import "strconv"

func solutions(start int, end int) string {
    result := ""

    for number := start; number < end; number++ {
        result += " " + strconv.Itoa(number)
    }

    return result
}

func fractiles(K int, C int, S int) string {
    if K == 1 {
        return " 1" // Bcz always the same tile + S always >= 1
    }

    if C == 1 { // Need to check all the tiles
        if S < K {
            return " IMPOSSIBLE"
        } else {
            return solutions(1, K + 1)
        }
    }

    if S < K - 1 { // Need to check the first K tiles, except the first one (tile #2 allows to know what is in the tile #1)
        return " IMPOSSIBLE"
    } else {
        return solutions(2, K + 1)
    }
}

func main() {
    var T int
    fmt.Scan(&T)

    for i := 1; i <= T; i++ {
        var K int
        var C int
        var S int
        fmt.Scan(&K)
        fmt.Scan(&C)
        fmt.Scan(&S)

        result := fractiles(K, C, S)

        fmt.Printf("Case #%d:%s\n", i, result)
    }
}
