package main

import "fmt"
import "strconv"
import "math/big"

func parse_number_base(jamcoin string, base int) int64 {
    bi := big.NewInt(0)
    _, is_success := bi.SetString(jamcoin, base)

    if !is_success {
        fmt.Println("ERROR Parse number")
        return -1
    }

    return bi.Int64()
}

func format_number_base(number int64, base int) string {
    return strconv.FormatInt(number, base)
}

func check_jamcoin(jamcoin string, N int) (bool, string) {
    if len(jamcoin) != N || jamcoin[0] != '1' || jamcoin[len(jamcoin) - 1] != '1' {
        return false, ""
    }

    result := ""

    for base := 2; base <= 10; base++ {
        number := parse_number_base(jamcoin, base)

        div := int64(2)
        for number % div != 0 {
            div += 1
            if div >= 100 {
                return false, ""
            }
        }
        result += " " + strconv.Itoa(int(div))
    }

    return true, result
}

func pow_int(number int, pow int) int64 {
    result := int64(number)

    for i := 0; i < pow; i++ {
        result *= int64(number)
    }

    return result
}

func generate_jamcoin(N int, J int) {
    number := pow_int(2, N - 2) + 1

    for j := 0; j < J; {
        jamcoin := format_number_base(number, 2)

        is_jamcoin, result := check_jamcoin(jamcoin, N)
        if is_jamcoin {
            fmt.Println(jamcoin + result)
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
