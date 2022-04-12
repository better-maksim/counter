package counter

type Counter interface {
	Add(num uint64)
	Read() uint64
}
