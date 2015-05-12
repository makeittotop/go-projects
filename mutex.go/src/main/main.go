package main

import (
    "fmt"
    "sync"
    "time"
)
func main() {
    m := new(sync.Mutex)
    // Serializing the goroutines through the mutex usage.
    for i := 0; i < 10; i++ {
        go func(i int) {
            m.Lock()
            fmt.Println(i, "Begin")
            time.Sleep(time.Second)
            fmt.Println(i, "End")
            m.Unlock()
        }(i)
    }

    var input string
    fmt.Scanln(&input)
}