package mylog

type Tags map[string]string
type Fields map[string]FieldValue
type Message string

type Logger struct {
	tags     Tags
	fields   Fields
	receiver Receiver
}

func (logger *Logger) Info(message Message) {
	logger.receiver.Info(logger.tags, logger.fields, message)
}

func (logger *Logger) Debug(message Message) {
	logger.receiver.Debug(logger.tags, logger.fields, message)
}

func (logger *Logger) Error(message Message) error {
	return logger.receiver.Error(logger.tags, logger.fields, message)
}

// TODO: more levels

func (logger *Logger) WithTag(name, value string) *Logger {
	logger.tags[name] = value
	return logger
}

func (logger *Logger) WithFields(fields ...*Field) *Logger {
	for _, field := range fields {
		logger.fields[field.Name] = field.Value
	}
	return logger
}

func cloneTags(raw Tags) Tags {
	res := make(Tags)
	for k, v := range raw {
		res[k] = v
	}
	return res
}

func cloneFields(raw Fields) Fields {
	// 这里不需要对 FieldValue 进行特殊clone，因为这些值不会在WithXXX时不会被修改（最多同一个key会被覆盖）。
	res := make(Fields)
	for k, v := range raw {
		res[k] = v
	}
	return res
}

func (logger *Logger) CloneLogger() *Logger {
	return &Logger{
		tags:     cloneTags(logger.tags),
		fields:   cloneFields(logger.fields),
		receiver: logger.receiver,
	}
}
