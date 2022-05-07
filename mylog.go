package mylog

var _logger *Logger

func Init(receiver Receiver) {
	_logger = &Logger{
		tags:     nil,
		fields:   nil,
		receiver: receiver,
	}
}

func CloneLogger() *Logger {
	return _logger.CloneLogger()
}
