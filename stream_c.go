package fuego

// This file is really an extension of stream.go.
//
// It contains the "concurrent" methods of Stream.
//
// The reason for splitting is two-fold:
// - convenient separation
// - help reduce the growth of stream.go

import "sync"

// PanicInvalidConcurrencyLevel signifies that the Stream is missing a channel.
const PanicInvalidConcurrencyLevel = "stream concurrency must be 1 or more"

// ForEachC is a concurrent wrapper of ForEach.
//
// The level of concurrency is set by the last call made to method
// Concurrent.
//
// See 'ForEach' for full details.
//
// Note that this method consumes the stream orderly but does NOT preserve
// order of output.
func (s Stream) ForEachC(consumer Consumer) {
	s.concurrentDo(func() { s.ForEach(consumer) })
}

// concurrentDo is a wrapper for concurrently running methods
// of Stream that are not concurrent.
//
// The level of concurrency is set by method 'Concurrent'.
//
// This function concurrently consumes from the stream's channel.
//
// Note that this method consumes the stream orderly but does NOT preserve
// order of output.
func (s Stream) concurrentDo(f func()) {
	s.panicIfInvalidConcurrency()

	var wg sync.WaitGroup

	for i := 1; i <= s.concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}

	wg.Wait()
}
