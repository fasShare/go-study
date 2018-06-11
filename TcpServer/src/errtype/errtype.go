package errtype

type StringError struct {
	str string
}

func (Err StringError) Error() string {
	return Err.str
}

func BuildStringError(val string) (err StringError) {
	err.str = val

	return err
}
