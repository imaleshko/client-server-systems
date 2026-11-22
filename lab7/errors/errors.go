package errors

type Error struct {
	err string
}

func (e *Error) Error() string {
	return e.err
}

type WrappedError struct {
	context string
	err     error
}

func (e *WrappedError) Error() string {
	return e.context + ": " + e.err.Error()
}

func (e *WrappedError) Unwrap() error {
	return e.err
}

func New(text string) error {
	return &Error{err: text}
}

func Wrap(context string, err error) error {
	if err == nil {
		return nil
	}
	return &WrappedError{context: context, err: err}
}

func unwrap(err error) error {
	w, ok := err.(*WrappedError)
	if !ok {
		return nil
	}
	return w.err
}

func Is(err, target error) bool {
	if target == nil {
		return err == nil
	}

	for ; err != nil; err = unwrap(err) {
		if err == target {
			return true
		}
	}

	return false
}

func As[T error](err error, target *T) bool {
	if target == nil {
		panic("target не може бути nil")
	}

	for ; err != nil; err = unwrap(err) {
		if v, ok := err.(T); ok {
			*target = v
			return true
		}
	}

	return false
}
