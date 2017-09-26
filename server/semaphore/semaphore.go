package semaphore

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

// simple semaphore implementation
type Semaphore chan struct{}

func (s Semaphore) Release() {
	<-s
}

func (s Semaphore) Wait() {
	select {
	case s <- struct{}{}:
	}
}