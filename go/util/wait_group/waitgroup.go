package wait_group

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

/*

w = WaitGroupWrapper{}
w.Wrap(job1)
w.Wrap(job2)

w.Wait()

 */
