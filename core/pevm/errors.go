package pevm

import (
	"fmt"
	"reflect"
)

type ErrBlocking struct {
	TxIdx int32
	err   error
}

func NewErrBlocking(txIdx int32, err error) ErrBlocking {
	return ErrBlocking{
		TxIdx: txIdx,
		err:   err,
	}
}

func (e ErrBlocking) Error() string {
	return fmt.Sprintf("transaction %d blocking: %v", e.TxIdx, e.err)
}

func (e ErrBlocking) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

type ErrExecution struct {
	TxIdx int32
	err   error
}

func NewErrExecution(txIdx int32, err error) ErrExecution {
	return ErrExecution{
		TxIdx: txIdx,
		err:   err,
	}
}

func (e ErrExecution) Error() string {
	return fmt.Sprintf("transaction %d execution error: %v", e.TxIdx, e.err)
}

func (e ErrExecution) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}
