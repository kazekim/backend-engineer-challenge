package grwrouters

import "sync"

func (r *defaultRouter) Start() error {

	var wg sync.WaitGroup
	wg.Add(1)

	wg.Wait()
	return nil
}
