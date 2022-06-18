package main

import (
	"fmt"

	"github.com/diegolopes98/go-data-structs/lists"
	singly "github.com/diegolopes98/go-data-structs/lists/sll"
)

func main() {
	sll := singly.NewSLL[int]()
	sll.Push(10).Push(20).Push(30).Push(40)
	singly.ForEach(
		singly.Filter(singly.Map(sll, func(v int) int {
			return v / 10
		}), func(v int) bool {
			return v%2 == 0
		}),
		func(n lists.Node[int]) {
			fmt.Println(n.GetValue())
		},
	)
}
