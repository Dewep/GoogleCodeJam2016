package main

import "fmt"

func sheep(N int) int {
    if N <= 0 {
        return -1
    }

    set := make(map[int]bool)
    for i := 0; i <= 9; i++ {
        set[i] = true
    }

    count := 0
    for len(set) > 0 {
        count += 1
        number := count * N

        for number != 0 {
            digit := number % 10
            delete(set, digit)
            number /= 10
        }
    }

    return count * N
}

func main() {
    var T int
    fmt.Scan(&T)

    for i := 1; i <= T; i++ {
        var N int
        fmt.Scan(&N)

        count := sheep(N)

        if count <= -1 {
            fmt.Printf("Case #%d: INSOMNIA\n", i)
        } else {
            fmt.Printf("Case #%d: %d\n", i, count)
        }
    }
}
