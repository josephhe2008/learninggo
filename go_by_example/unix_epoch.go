package main
import (
    "fmt"
    "time"
)
func main() {

	// unix epoch 1 January 1970
    now := time.Now()
    fmt.Println(now)
    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}