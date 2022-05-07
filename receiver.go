package mylog

type Receiver interface {
	Info(tags Tags, fields Fields, message Message)
	Debug(tags Tags, fields Fields, message Message)
	Error(tags Tags, fields Fields, message Message) error
}
