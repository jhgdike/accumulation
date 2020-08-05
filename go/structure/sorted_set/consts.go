package zset

import "errors"

const (
	ZSKIPLIST_P = 0.25

	ZADD_NONE    = 0
	ZADD_INCR    = 1 << iota /* Increment the score instead of setting it. */
	ZADD_NX                  /* Don't touch elements not already existing. */
	ZADD_XX                  /* Only touch elements already existing. */
	ZADD_NOP                 /* Operation not performed because of conditionals.*/
	ZADD_NAN                 /* Only touch elements already existing. */
	ZADD_ADDED               /* The element was new and was added. */
	ZADD_UPDATED             /* The element already existed, score updated. */

	/* Flags only used by the ZADD command but not by zsetAdd() API: */
	ZADD_CH = 1 << 16 /* Return num of elements added or updated. */

)

var (
	notFoundError = errors.New("not found")
)
