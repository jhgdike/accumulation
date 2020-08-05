package zset

import (
	"math/rand"
	"time"
)

const (
	ZSkipListMaxLevel = 32
)

func init() {
	rand.Seed(time.Now().Unix())
}

type zSkipList struct {
	header, tail *zSkipListNode
	length       int
	level        int
}

func NewZSkipList() *zSkipList {
	return &zSkipList{
		level:  1,
		length: 0,
		header: zslCreateNode(ZSkipListMaxLevel, 0, ""),
		tail:   nil,
	}
}

func (zsl *zSkipList) InsertNode(key string, score float64) *zSkipListNode {
	var x *zSkipListNode
	update := make([]*zSkipListNode, ZSkipListMaxLevel)
	rank := make([]int, ZSkipListMaxLevel)

	x = zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		if i == (zsl.level - 1) {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.level[i].forward != nil && (
			x.level[i].forward.score < score || (
				x.level[i].forward.score == score && x.level[i].forward.key < key)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}

		update[i] = x
	}
	/* we assume the element is not already inside, since we allow duplicated
	 * scores, reinserting the same element should never happen since the
	 * caller of zslInsert() should test in the hash table if the element is
	 * already inside or not. */

	level := zslRandomLevel()
	if level > zsl.level {
		for i := zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = int(zsl.length)
		}
		zsl.level = level
	}
	x = zslCreateNode(level, score, key)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = rank[0] - rank[i] + 1
	}

	/* increment span for untouched levels */
	for i := level; i < zsl.level; i++ {
		update[i].level[i].span += 1
	}

	if update[0] != zsl.header {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}
	zsl.length += 1
	return x
}

func (zsl *zSkipList) Update(key string, curScore, newScore float64) *zSkipListNode {
	update := make([]*zSkipListNode, ZSkipListMaxLevel)

	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (
			x.level[i].forward.score < curScore || (
				x.level[i].forward.score == curScore && x.level[i].forward.key < key)) {
			x = x.level[i].forward
		}

		update[i] = x
	}

	x = x.level[0].forward
	if x == nil || curScore != x.score || x.key != key {
		panic("logic error, zNode not found")
	}
	if (x.backward == nil || x.backward.score < newScore) && (x.level[0].forward == nil || x.level[0].forward.score > newScore) {
		x.score = newScore
		return x
	}

	zsl.DeleteNode(x, update)
	return zsl.InsertNode(key, newScore)
}

func (zsl *zSkipList) Delete(key string, score float64) *zSkipListNode {
	update := make([]*zSkipListNode, ZSkipListMaxLevel)

	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (
			x.level[i].forward.score < score || (
				x.level[i].forward.score == score && x.level[i].forward.key < key)) {
			x = x.level[i].forward
		}

		update[i] = x
	}

	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		zsl.DeleteNode(x, update)
		return x
	}
	return nil
}

func (zsl *zSkipList) DeleteNode(x *zSkipListNode, update []*zSkipListNode) {
	for i := 0; i < zsl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		zsl.tail = x.backward
	}
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level -= 1
	}
	zsl.length -= 1
}

func (zsl *zSkipList) Rank(key string, score float64) int {
	x := zsl.header
	var rank int
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (
			x.level[i].forward.score < score || (
				x.level[i].forward.score == score && x.level[i].forward.key < key)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}
	}
	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		return rank
	}
	return -1
}

func (zsl *zSkipList) Range(start, end int) (res []*zSkipListNode) {
	var rank int
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && x.level[i].span+rank <= start + 1 {
			rank += x.level[i].span
			x = x.level[i].forward
		}
	}
	if rank < start {
		return
	} else {
		for i := 0; i < end-start; i++ {
			if x == nil {
				break
			}
			res = append(res, x)
			x = x.level[0].forward
		}
	}
	return
}

type zSkipListNode struct {
	key      string
	score    float64
	backward *zSkipListNode // 回退，每一层都有自己的前置节点
	level    []nodeLevel
}

type nodeLevel struct {
	forward *zSkipListNode // 下一跳节点
	span    int            // 到下一跳节点的距离
}

func zslCreateNode(level int, score float64, key string) *zSkipListNode {
	return &zSkipListNode{
		key:   key,
		score: score,
		level: make([]nodeLevel, level),
	}
}

func zslRandomLevel() int {
	level := 1
	for float64(rand.Int()&0xffff) < (ZSKIPLIST_P * 0xffff) {
		level += 1
	}
	if level < ZSkipListMaxLevel {
		return level
	} else {
		return ZSkipListMaxLevel
	}
}
