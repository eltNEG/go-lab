package main

import (
	"fmt"

	example_simple "github.com/eltneg/lab/learn-grpc/src"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "kdsjdsdsa dsad",
		SimpleList: []int32{1, 3, 4, 5, 3},
	}
	fmt.Println(sm)
	fmt.Println(sm.GetId())
}
