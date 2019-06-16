package util

import (
	"fmt"
	"testing"
)

func TestBytesToInt64(t *testing.T) {
	fmt.Println(BytesToInt64([]byte{0, 0,0,0,0,0,0,1}))
}

func TestInt64ToBytes(t *testing.T) {
	fmt.Println(Int64ToBytes(1024))
}
