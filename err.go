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
	once       sync.Once // 用于确保通道只被关闭一次

}

func New(callback ErrorHandleFunc, shouldWait bool) *Error {
	return &Error{
		errCh:    make(chan error),
		callback: callback,
		// 错误处理
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
