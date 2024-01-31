// Code for program Q (The Server)
package server

import (
	"math"
)

type RPCServer int

//procedure1 takes one integer argument

type IncrementArgs struct {
	X int
}

//procedure2 takes  2 arguments , a float and an integer

type AdditionArgs struct {
	X float64
	Y int
}

//procedure3 takes one string argument
type ReverseStringArgs struct{
	Word string
}


//in golang rpc methods this signature
// func (t *T) MethodName(argType T1, replyType *T2) error
func (s *RPCServer) Procedure1(args *IncrementArgs, result *int) error {

	*result = args.X + 1
	return nil

}

func (s *RPCServer) Procedure2(args *AdditionArgs, result *float64) error {

	*result = math.Pow(args.X, 2) + float64(args.Y)
	return nil

}

func ( s * RPCServer) Procedure3(args *ReverseStringArgs, result *string) error{
	//convert to runes
	rns:= []rune(args.Word)
	
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        // swap the letters of the string, 
        // eg 1st with the last and so on. 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    // return the reversed string. 
    *result=string(rns) 
	return nil
	
}