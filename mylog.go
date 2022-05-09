// Package mylog 是一个相对通用的日志入口，可以对接其他日志库使用。

package mylog

var _logger = &Logger{
	tags:     nil,
	fields:   nil,
	receiver: &dummyReceiver{},
}

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
