package common

import (
	"fmt"
)

func HandlePanic(h Handler, logger Logger) error {
	return handlePanic(recover(), h, logger)
}

func (hf hFunc) Filter(err error, r interface{}) error {
	return hf(err, r)
}

func (lf loggerFunc) Ef(format string, a ...interface{}) {
	lf(format, a)
}

func HandlePanicFunc(hf func(err error, r interface{}) error, logger func(format string, a ...interface{})) error {
	var f Handler
	if hf != nil {
		f = hFunc(hf)
	}
	var l Logger
	if logger != nil {
		l = loggerFunc(logger)

	}
	return handlePanic(recover(), f, l)
}

func handlePanic(r interface{}, h Handler, logger Logger) error {
	if r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		if h != nil {
			err = h.Filter(err, r)
		}
		if err != nil && logger != nil {
			logger.Ef("panic err %+v", err)
		}
		return err
	}
	return nil
}
