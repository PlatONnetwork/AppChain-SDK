package pevm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/PlatONnetwork/PlatON-Go/common"
)

var (
	ErrInconsistentRead       = InconsistentReadError{}
	ErrInvalidMemoryValueType = InvalidMemoryValueTypeError{}
	ErrSelfDestructedAccount  = SelfDestructedAccountError{}

	ErrRetry                = RetryError{}
	ErrFallbackToSequential = FallbackToSequentialError{}
)

type StorageError struct {
	Reason string
}

func (e StorageError) Error() string {
	return fmt.Sprintf("Failed to reading memory from storage: %s", e.Reason)
}

func (e StorageError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type BlockingError struct {
	Addr  common.Address
	TxIdx int32
}

func (e BlockingError) Error() string {
	return fmt.Sprintf("Read of memory location(%s) is blocked by tx #%d", e.Addr.Hex(), e.TxIdx)
}

func (e BlockingError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type InconsistentReadError struct{}

func (e InconsistentReadError) Error() string {
	return "Inconsistent read"
}

func (e InconsistentReadError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type InvalidNonceError struct {
	TxIdx int32
}

func (e InvalidNonceError) Error() string {
	return fmt.Sprintf("Tx %d has invalid nonce", e.TxIdx)
}

func (e InvalidNonceError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type SelfDestructedAccountError struct{}

func (e SelfDestructedAccountError) Error() string {
	return "Tried to read self-destructed account"
}

func (e SelfDestructedAccountError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type InvalidBytecodeError struct {
	Source error
}

func (e InvalidBytecodeError) Error() string {
	return fmt.Sprintf("Invalid bytecode: %v", e.Source)
}

func (e InvalidBytecodeError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type InvalidMemoryValueTypeError struct{}

func (e InvalidMemoryValueTypeError) Error() string {
	return "Invalid type of stored memory value"
}

func (e InvalidMemoryValueTypeError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type RetryError struct{}

func (e RetryError) Error() string {
	return "Retry"
}

func (e RetryError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type FallbackToSequentialError struct{}

func (e FallbackToSequentialError) Error() string {
	return "Fallback to sequential"
}

func (e FallbackToSequentialError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type ExecutionBlockingError struct {
	TxIdx int32
}

func (e ExecutionBlockingError) Error() string {
	return fmt.Sprintf("Tx #%d blocked", e.TxIdx)
}

func (e ExecutionBlockingError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type ExecutionError struct {
	err error
}

func (e ExecutionError) Error() string {
	return fmt.Sprintf("Execution error: %v", e.err)
}

func (e ExecutionError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

func ToVmExecutionError(err error) error {
	switch {
	case errors.Is(err, ErrInconsistentRead):
		return ErrRetry
	case errors.Is(err, ErrSelfDestructedAccount):
		return ErrFallbackToSequential
	case errors.Is(err, BlockingError{}):
		blockingErr := err.(BlockingError)
		return ExecutionBlockingError{blockingErr.TxIdx}
	default:
		return ExecutionError{err}
	}
}
