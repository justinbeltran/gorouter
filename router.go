package main

import "fmt"
import zmq "github.com/alecthomas/gozmq"

func main() {
	context, _ := zmq.NewContext()
	pull, _ := context.NewSocket(zmq.PULL)
	push, _ := context.NewSocket(zmq.PUSH)
	pull.Bind("tcp://127.0.0.1:5000")
	push.Connect("tcp://127.0.0.1:6000")
	c := make(chan string)
	go recv(c, pull)
	go send(c, push)
	select {} //block forever
}

func recv(c chan string, socket *zmq.Socket) {
	for {
		msg, _ := socket.Recv(0)
		fmt.Println("Pulled msg: ", string(msg))
		c <- string(msg)
	}
}

func send(c chan string, socket *zmq.Socket) {
	for {
		msg := <-c
		socket.Send([]byte(msg), 0)
		fmt.Println("Pushed msg: ", msg)
	}
}
