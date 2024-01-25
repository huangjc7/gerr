package gerr

import (
	"sync"
)

type ErrorHandleFunc func(err error)

type Error struct {
	errCh      chan error
	wg         sync.WaitGroup
	callback   ErrorHandleFunc
	shouldWait bool
	once       sync.Once // Used to ensure that the channel is only closed once

}

func New(callback ErrorHandleFunc, shouldWait bool) *Error {
	// When shouldWait is true, waitGroup will be used to facilitate the scenario where the goroutine has not completed execution when the function exits.
	// When shouldWait is false, waitGroup will not be used to continue receiving errors from the error channel.
	// Note: When shouldWait is false, there is no need to call the Close method
	return &Error{
		errCh:      make(chan error),
		callback:   callback,
		shouldWait: shouldWait,
	}
}

func (e *Error) CatchError(err error) {
	if err != nil {
		e.errCh <- err
	}
}

func (e *Error) Receive() {
	if e.shouldWait {
		e.wg.Add(1)
	}

	go func() {
		if e.shouldWait {
			defer e.wg.Done()
		}

		for err := range e.errCh {
			if e.callback != nil {
				e.callback(err)
			}
		}
	}()

}

func (e *Error) Close() {
	e.once.Do(func() {
		close(e.errCh)
	})

	if e.shouldWait {
		e.wg.Wait()
	}
}
