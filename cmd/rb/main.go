package main

import (
	"fmt"

	rb "github.com/dorrella/matrix-curses/pkg/ringbuff"
)

func print(r *rb.RingBuff) {
	for i := 0; i < r.GetSize(); i++ {
		line, ok := r.Get(i)
		if !ok {
			panic("bad")
		}

		fmt.Printf("%s ", line)
	}
	fmt.Println()
}

func main() {
	buff := rb.NewRingBuff(3)
	fmt.Println(buff.IsEmpty())
	fmt.Println(buff.GetSize())
	fmt.Println()

	buff.Add("1")
	fmt.Println(buff.IsEmpty())
	fmt.Println(buff.GetSize())
	fmt.Println()
	print(buff)

	buff.Add("2")
	fmt.Println(buff.GetSize())
	fmt.Println()
	print(buff)

	buff.Add("3")
	fmt.Println(buff.GetSize())
	fmt.Println()
	print(buff)

	buff.Add("4")
	fmt.Println(buff.GetSize())
	fmt.Println()
	print(buff)

	buff.Add("5")
	fmt.Println(buff.GetSize())
	fmt.Println()
	print(buff)
}
