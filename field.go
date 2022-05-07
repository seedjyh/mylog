package mylog

type FieldValue interface{}

type Field struct {
	Name  string
	Value FieldValue
}

func String(name string, value string) *Field {
	return &Field{
		Name:  name,
		Value: value,
	}
}

func Int(name string, value int) *Field {
	return &Field{
		Name:  name,
		Value: value,
	}
}

func Float(name string, value float32) *Field {
	return &Field{
		Name:  name,
		Value: value,
	}
}

func Error(name string, value error) *Field {
	return &Field{
		Name:  name,
		Value: value.Error(),
	}
}
