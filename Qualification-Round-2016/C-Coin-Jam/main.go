package main

import "fmt"
import "strconv"
import "math"

func is_prime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n & 1 == 0 {
        return false
    }
    s := int(math.Sqrt(float64(n)))
    for i := 3; i <= s; i += 2 {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func parse_number_base(jamcoin string, base int) int {
    value, err := strconv.ParseInt(jamcoin, base, 64)

    if err != nil {
        return -1
    }

    return int(value)
}

func format_number_base(number int, base int) string {
    return strconv.FormatInt(int64(number), base)
}

func is_jamcoin(jamcoin string, N int) bool {
    if len(jamcoin) != N || jamcoin[0] != '1' || jamcoin[len(jamcoin) - 1] != '1' {
        return false
    }

    for i := 2; i <= 10; i++ {
        number := parse_number_base(jamcoin, i)

        if number == -1 || is_prime(number) {
            return false
        }
    }

    return true
}

func get_min_number_jamcoin(N int, J int) int {
    var jamcoin string

    for n := 0; n < N; n++ {
        if n == 0 || n == N - 1 {
            jamcoin += "1"
        } else {
            jamcoin += "0"
        }
    }

    return parse_number_base(jamcoin, 2)
}

func generate_jamcoin(N int, J int) {
    number := get_min_number_jamcoin(N, J)

    for j := 0; j < J; {
        jamcoin := format_number_base(number, 2)

        if is_jamcoin(jamcoin, N) {
            fmt.Println(jamcoin)
            j += 1
        }

        number += 1
    }
}

func main() {
    var T int
    fmt.Scan(&T)

    for i := 1; i <= T; i++ {
        var N int
        var J int
        fmt.Scan(&N)
        fmt.Scan(&J)

        fmt.Printf("Case #%d:\n", i)
        generate_jamcoin(N, J)
    }
}
