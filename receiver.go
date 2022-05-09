package mylog

type Receiver interface {
	Info(tags Tags, fields Fields, message Message)
	Debug(tags Tags, fields Fields, message Message)
	Error(tags Tags, fields Fields, message Message) error
}

type dummyReceiver struct{}

func (r *dummyReceiver) Info(tags Tags, fields Fields, message Message)        {}
func (r *dummyReceiver) Debug(tags Tags, fields Fields, message Message)       {}
func (r *dummyReceiver) Error(tags Tags, fields Fields, message Message) error { return nil }
