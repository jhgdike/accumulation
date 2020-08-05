package zset

type ZSet struct {
	dict      map[string]*zSkipListNode
	zSkipList *zSkipList
}

func NewZSet() *ZSet {
	return &ZSet{
		dict:      make(map[string]*zSkipListNode),
		zSkipList: NewZSkipList(),
	}
}

func (z *ZSet) Add(key string, score float64) {
	flag := ZADD_NX
	z.ZAdd(key, score, &flag, nil)
}

func (z *ZSet) Update(key string, score float64) {
	flag := ZADD_XX
	z.ZAdd(key, score, &flag, nil)
}

func (z *ZSet) ZAdd(key string, score float64, flags *int, newScore *float64) int {
	// Check vars
	incr := (*flags & ZADD_INCR) != 0
	nx := (*flags & ZADD_NX) != 0
	xx := (*flags & ZADD_XX) != 0
	*flags = 0

	var curScore float64
	var zNode *zSkipListNode

	zNode, ok := z.dict[key]
	if ok && zNode != nil {
		if nx {
			*flags |= ZADD_NOP
			return 1
		}
		curScore = zNode.score

		//incr
		if incr {
			score += curScore
			if newScore != nil {
				*newScore = score
			}
		}

		// remove and re-insert when score changed.
		if score != curScore {
			node := z.zSkipList.Update(key, curScore, score)
			z.dict[key] = node
			*flags |= ZADD_UPDATED
		}
		return 1
	} else if !xx {
		node := z.zSkipList.InsertNode(key, score)
		z.dict[key] = node
		*flags |= ZADD_ADDED
		if newScore != nil {
			*newScore = score
		}
		return 1
	} else {
		*flags |= ZADD_NOP
		return 1
	}
}

func (z *ZSet) Del(key string) *float64 {
	if zNode, ok := z.dict[key]; ok {
		delete(z.dict, key)
		z.zSkipList.Delete(key, zNode.score)
		return &zNode.score
	} else {
		return nil
	}
}

func (z *ZSet) Rank(key string) int {
	if zNode, ok := z.dict[key]; ok {
		return z.zSkipList.Rank(key, zNode.score)
	} else {
		return -1
	}
}

func (z *ZSet) RevRank(key string) int {
	rank := z.Rank(key)
	if rank != -1 {
		rank = z.zSkipList.length - rank - 1
	}
	return rank
}

func (z *ZSet) IncrBy(key string, increment float64) (score float64) {
	flag := ZADD_INCR
	z.ZAdd(key, increment, &flag, &score)
	return
}

func (z *ZSet) Range(start, end int) (res []RangeResp) {
	if start >= end {
		return
	}
	list := z.zSkipList.Range(start, end)
	for _, l := range list {
		res = append(res, RangeResp{Key: l.key, Score: l.score})
	}
	return
}

type RangeResp struct {
	Key   string
	Score float64
}
