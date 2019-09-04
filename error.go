package envconf

import "fmt"

type Error struct {
	Inner     error
	Message   string
	FieldName string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s. field=%s inner_err=%s", e.Message, e.FieldName, e.Inner)
}
