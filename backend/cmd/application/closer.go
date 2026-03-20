package application

import (
	"log"
	"sync"
)

type closers struct {
	closers     []func() error
	closersLock sync.Mutex
}

func (a *closers) AddCloser(closer func() error) {
	a.closersLock.Lock()
	defer a.closersLock.Unlock()

	a.closers = append(a.closers, closer)
}

func (a *closers) Close() {
	a.closersLock.Lock()
	defer a.closersLock.Unlock()

	for _, closer := range a.closers {
		if err := closer(); err != nil {
			log.Printf("closer: %s", err)
		}
	}
}
