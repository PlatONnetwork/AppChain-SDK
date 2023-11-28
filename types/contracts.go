package types

var (
	UndefinedError = []byte("undefined error")
)

type RevertError struct {
	error
	ReturnData []byte
}

func (e *RevertError) Error() string {
	return string(e.ReturnData)
}

func NewRevertError(data string) *RevertError {
	return &RevertError{
		ReturnData: []byte(data),
	}
}
