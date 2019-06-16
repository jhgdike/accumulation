package byte_order

import (
	"bytes"
	"encoding/binary"
)

//func Int64ToBytes(num int64) []byte {
//	buf := make([]byte, 8)
//	binary.PutVarint(buf, num)
//	return buf
//}
//
//func BytesToInt64(bytes []byte) int64 {
//	ans, _ := binary.Varint(bytes)
//	return ans
//}

//func Int64ToBytes(i int64) []byte {
//	var buf = make([]byte, 8)
//	binary.BigEndian.PutUint64(buf, uint64(i))
//	return buf
//}

//func BytesToInt64(buf []byte) int64 {
//	return int64(binary.BigEndian.Uint64(buf))
//}

// 大端字节序
// 1024  -->  [0 0 0 0 0 0 4 0]  (十六进制: Ox400，二进制: 100 0000 0000)
func Int64ToBytes(i int64) []byte {
	b := make([]byte, 8)
	v := uint64(i)

	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b

}

// [0 0 0 0 0 0 4 0]  (十六进制: Ox400，二进制: 100 0000 0000)   -->  1024
func BytesToInt64(b []byte) int64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return int64(uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56)
}

func Int64ToBytes2(i int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	return bytesBuffer.Bytes()
}

func BytesToInt642(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

// 小端字节序
func LInt64ToBytes(i int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, i)
	return bytesBuffer.Bytes()
}

func LBytesToInt642(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}
