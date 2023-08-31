package main
import "fmt"
func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

	//messages <- msg
	//mc := <-messages
	//fmt.Println(mc)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

// An unbuffered channel is a channel that needs a receiver as soon as a 
// message is emitted to the channel. 
// The goroutine is blocked after sending the message since no receivers are 
// yet ready. Accoring to the spec, if the capacity is zero or absent, the 
// channel is unbuffered and the communication succeeds only when both a
// sender and receiver are ready.
// If the channel is unbuffered, the sender blocks until the receiver has
// received the value.