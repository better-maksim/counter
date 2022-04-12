package counter

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Acc266(counter Counter) uint64 {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter Counter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter Counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter Counter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}
	wg.Wait()
	return counter.Read()
}

func TestAtomicCounter(t *testing.T) {
	counter := NewAtomicCounter()
	value := Acc266(counter)
	assert.Equal(t, value, uint64(66))
}

func TestCASCounter(t *testing.T) {
	counter := NewCASCounter()
	value := Acc266(counter)
	assert.Equal(t, value, uint64(66))
}

func TestChannelCounter(t *testing.T) {
	counter := NewChannelCounter()
	value := Acc266(counter)
	assert.Equal(t, value, uint64(66))
}

func TestMutexCounter(t *testing.T) {
	counter := NewMutexCounter()
	value := Acc266(counter)
	assert.Equal(t, value, uint64(66))
}
