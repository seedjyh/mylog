package main

import (
	"errors"
	"github.com/seedjyh/mylog"
	"github.com/seedjyh/mylog/receiver/seelog"
)

func main() {
	mylog.Init(seelog.NewReceiverFromConfigAsFile("example/easy-json/seelog.xml"))
	inheritedFields()
	independentFields()
	coveredFields()
	tagsAndFields()
}

func inheritedFields() {
	logger := mylog.CloneLogger()

	// fields: {"age":10}
	logger.WithFields(mylog.Int("age", 10)).Info("Only age.")
	// fields: {"age":10, "gender":"male"}
	logger.WithFields(mylog.String("gender", "male")).Info("Age & gender.")
}

func independentFields() {
	logger := mylog.CloneLogger()

	// fields: {"age":10}
	logger.CloneLogger().WithFields(mylog.Int("age", 10)).Info("Only age.")
	// fields: {"gender":"male"}
	logger.CloneLogger().WithFields(mylog.String("gender", "male")).Info("Only gender.")
}

func coveredFields() {
	logger := mylog.CloneLogger()

	// fields: {"age":10}
	logger.WithFields(mylog.Int("age", 10)).Info("Ten.")
	// fields: {"age":20}
	logger.WithFields(mylog.Int("age", 20)).Info("Twenty.")
}

func tagsAndFields() {
	_ = mylog.CloneLogger().
		WithTag("user_id", "user_123").      // tag 的值一定是字符串
		WithTag("session_id", "2147480001"). // tag 的值一定是字符串
		WithFields(mylog.Error("the_error", errors.New("some-error"))).
		Error("Here's some error.")
}
