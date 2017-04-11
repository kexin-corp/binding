package binding

var logger BLogger

type BLogger interface {
	Errorf(format string, args ...interface{})
}

func SetLogger(l BLogger) {
	logger = l
}

func Errorf(format string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(format, args)
	}
}