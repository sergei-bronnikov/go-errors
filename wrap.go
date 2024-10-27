package errors

type wrapError struct {
	err          error
	wrappedError error
}

func (w *wrapError) Error() string {
	return w.err.Error()
}

func (w *wrapError) Message() string {
	switch e := w.err.(type) {
	case interface{ Message() string }:
		return e.Message()
	default:
		return ""
	}
}

func (w *wrapError) Context() string {
	switch e := w.err.(type) {
	case interface{ Context() string }:
		return e.Context()
	default:
		return ""
	}
}

func (w *wrapError) Is(target error) bool {
	return w.Error() == target.Error()
}

func (w *wrapError) Unwrap() error {
	return w.wrappedError
}
