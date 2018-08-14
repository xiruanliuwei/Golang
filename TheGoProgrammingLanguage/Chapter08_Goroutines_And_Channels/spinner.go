
package main

import "fmt"
import "time"

func main() {
    go spinner(100 * time.Millisecond)

    const n = 45    // n type: int, %T
    fibN := fib(n)

    fmt.Printf("n type is: %T\n", n);

    fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {    // type Duration int64

    for {
        for _, r := range `-\|/` {    // raw string literals with back quotes
            fmt.Printf("\r%c", r)
            time.Sleep(delay)    // Sleep pauses the current goroutine
        }
    }

}

func fib(x int) int {
    if x < 2 {
        return x
    }

    return fib(x-1) + fib(x-2)
}
