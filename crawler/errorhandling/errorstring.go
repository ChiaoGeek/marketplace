package errorhandling

type Errorstring struct {
	S string
}

func (err *Errorstring) Error() string{
	return err.S
}

func Newerror(text string) error {
	return &Errorstring{S:text}
}

