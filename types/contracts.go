package types

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
