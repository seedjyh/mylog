package seelog

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/seedjyh/mylog"
)

type receiver struct {
	loggerInterface seelog.LoggerInterface
}

func init() {
	if err := seelog.RegisterCustomFormatter("Tags", createTagsFormatter); err != nil {
		panic(err)
	}
	if err := seelog.RegisterCustomFormatter("Fields", createFieldsFormatter); err != nil {
		panic(err)
	}
}

func NewReceiverFromConfigAsFile(configFile string) *receiver {
	if _logger, err := seelog.LoggerFromConfigAsFile(configFile); err != nil {
		panic(err)
	} else {
		return &receiver{
			loggerInterface: _logger,
		}
	}
}

type seelogContext struct {
	tags   mylog.Tags
	fields mylog.Fields
}

func createTagsFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		if logContext, ok := context.CustomContext().(*seelogContext); !ok {
			return "Broken Context!"
		} else {
			if buf, err := json.Marshal(logContext.tags); err != nil {
				return err.Error()
			} else {
				return string(buf)
			}
		}
	}
}

func createFieldsFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		if logContext, ok := context.CustomContext().(*seelogContext); !ok {
			return "Broken Context!"
		} else {
			if buf, err := json.Marshal(logContext.fields); err != nil {
				return err.Error()
			} else {
				return string(buf)
			}
		}
	}
}

func (r *receiver) Debug(tags mylog.Tags, fields mylog.Fields, message mylog.Message) {
	logger := r.withContext(&seelogContext{
		tags:   tags,
		fields: fields,
	})
	defer logger.Flush()
	logger.Debug(message)
}

func (r *receiver) Info(tags mylog.Tags, fields mylog.Fields, message mylog.Message) {
	logger := r.withContext(&seelogContext{
		tags:   tags,
		fields: fields,
	})
	defer logger.Flush()
	logger.Info(message)
}

func (r *receiver) Error(tags mylog.Tags, fields mylog.Fields, message mylog.Message) error {
	logger := r.withContext(&seelogContext{
		tags:   tags,
		fields: fields,
	})
	defer logger.Flush()
	return logger.Error(message)
}

func (r *receiver) withContext(c *seelogContext) seelog.LoggerInterface {
	if logger, err := seelog.CloneLogger(r.loggerInterface); err != nil {
		return r.loggerInterface // 降级
	} else {
		logger.SetContext(c)
		return logger
	}
}
