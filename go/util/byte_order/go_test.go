package byte_order

import (
	"fmt"
	"testing"
)

func TestBytesToInt64(t *testing.T) {
	fmt.Println(BytesToInt64([]byte{0, 0,0,0,0,0,4,0}))  // 1024
	fmt.Println(BytesToInt642([]byte{0, 0,0,0,0,0,4,0}))  // 1024
	fmt.Println(LBytesToInt642([]byte{0, 0,0,0,0,0,4,0}))  // 1125899906842624
}

func TestInt64ToBytes(t *testing.T) {
	fmt.Println(Int64ToBytes(1024))  // [0 0 0 0 0 0 4 0]
	fmt.Println(Int64ToBytes2(1024))  // [0 0 0 0 0 0 4 0]
	fmt.Println(LInt64ToBytes(1024))  // [0 4 0 0 0 0 0 0]
}
