package main

import "fmt"
import zmq "github.com/alecthomas/gozmq"

func main() {
	context, _ := zmq.NewContext()
	push, _ := context.NewSocket(zmq.PUSH)
	push.Connect("tcp://127.0.0.1:5000")
	pull, _ := context.NewSocket(zmq.PULL)
	pull.Bind("tcp://127.0.0.1:6000")

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("msg %d", i)
		push.Send([]byte(msg), 0)
		println("Sent", msg)
		forward, _ := pull.Recv(0)
		println("Received", string(forward))
	}
}
