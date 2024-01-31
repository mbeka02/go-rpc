package main

import (
	"fmt"
	"log"

	"net"
	"net/http"
	"net/rpc"

	"github.com/mbeka02/go-rpc/server"
)

// main function
func main() {
	// The server or program Q , check the server directory to see all 3 procedures
	s := new(server.RPCServer)
	rpc.Register(s)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)
	runProgramP()

}

// The 'Client or program P

func runProgramP() {
   //connect to the rpc server
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//Call procedure 1
	args := &server.IncrementArgs{X: 100}
	var result int
	err = client.Call("RPCServer.Procedure1", args, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" 1 + %v is %v \n", args.X, result)

	//Call procedure 2
	args2 := &server.AdditionArgs{X: 12.5, Y: 3}
	var result2 float64
	err = client.Call("RPCServer.Procedure2", args2, &result2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The result is : %v \n", result2)

	//Call Procedure 3
	args3 := &server.ReverseStringArgs{Word: "Anthony"}
	var result3 string
	err = client.Call("RPCServer.Procedure3", args3, &result3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The result is : %s \n", result3)

}
