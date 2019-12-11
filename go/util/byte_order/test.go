package main

import (
	"context"
	"fmt"
)

func main() {
	var ctx = context.Background()
	select {
	default:
		fmt.Println("aaa")
		case <- ctx.Done():
			fmt.Println("bbb")
			return
	}
	fmt.Println("ccc")
}
