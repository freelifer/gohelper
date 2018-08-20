package e

type Err interface {
	Code() int
	Error() string
}

type BaseErr struct {
	C int
	E string
}

func (e *BaseErr) Code() int {
	return e.C
}

func (e *BaseErr) Error() string {
	return e.E
}

func New(code int, msg string) Err {
	return &BaseErr{code, msg}
}

func NewInnerErr(msg string) Err {
	return &BaseErr{INNER_ERROR, msg}
}
