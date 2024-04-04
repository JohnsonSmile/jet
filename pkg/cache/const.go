package cache

type Size uint64

const (
	B Size = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)
