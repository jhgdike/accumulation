package zset

import (
	"fmt"
	"testing"
)

func TestNewZSet(t *testing.T) {
	z := NewZSet()
	z.Add("a", 1)
	z.Add("b", 2)
	z.Add("c", 3)

	fmt.Println(z.Range(0, 3))  // [{a 1} {b 2} {c 3}]
	fmt.Println(z.Range(0, 2))  // [{a 1} {b 2}]
	fmt.Println(z.Range(0, 1))  // [{a 1}]
	fmt.Println(z.Range(1, 3))  // [{b 2} {c 3}]
	fmt.Println(z.Range(1, 2))  // [{b 2}]
	fmt.Println(z.Range(1, 1))  // []

	fmt.Println(z.RevRank("a"))  // 2
	fmt.Println(z.RevRank("b"))  // 1
	fmt.Println(z.RevRank("c"))  // 0

	fmt.Println(z.Rank("a"))  // 0
	fmt.Println(z.Rank("b"))  // 1
	fmt.Println(z.Rank("c"))  // 2

	z.Update("b", 5)
	fmt.Println("\n zRank after update")
	fmt.Println(z.Rank("a"))  // 0
	fmt.Println(z.Rank("b"))  // 2
	fmt.Println(z.Rank("c"))  // 1
	fmt.Println(z.Range(0, 3))  // [{a 1} {c 3} {b 5}]

	z.IncrBy("c", 7)
	fmt.Println("\n zRank after incr")
	fmt.Println(z.Rank("a"))  // 0
	fmt.Println(z.Rank("b"))  // 2
	fmt.Println(z.Rank("c"))  // 1
	fmt.Println(z.Range(0, 3))  // [{a 1} {b 5} {c 10}]

	z.Del("b")
	fmt.Println("\n zRank after del")
	fmt.Println(z.Rank("a"))  // 0
	fmt.Println(z.Rank("c"))  // 1
	fmt.Println(z.Range(0, 3))  // [{a 1} {c 10}]

	t.Log()
}
