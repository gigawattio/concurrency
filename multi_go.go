package concurrency

import (
	"sync"

	"gigawatt.io/errorlib"
)

// MultiGo runs the passed `func() error` functions in parallel and then merges
// any resulting errors.
func MultiGo(funcs ...func() error) error {
	var (
		errs    = []error{}
		errLock sync.Mutex
		wg      sync.WaitGroup
	)

	for i, fn := range funcs {
		if fn != nil {
			wg.Add(1)
			go func(fn func() error, i int) {
				if err := fn(); err != nil {
					errLock.Lock()
					errs = append(errs, err)
					errLock.Unlock()
				}
				wg.Done()
			}(fn, i)
		}
	}

	wg.Wait()

	if err := errorlib.Merge(errs); err != nil {
		return err
	}
	return nil
}
