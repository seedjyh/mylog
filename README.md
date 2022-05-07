# mylog

mylog 是一个相对通用的日志入口，可以对接其他日志库使用。

## 快速上手

进入根目录并执行`go run example/easy-json/main.go`。

如果有如下输出就是执行成功。

```text
{"time":1651911648435952800,"lev":"Info","msg":"Only age.", "tags":{}, "fields":{"age":10}}
{"time":1651911648436514900,"lev":"Info","msg":"Age & gender.", "tags":{}, "fields":{"age":10,"gender":"male"}}
{"time":1651911648437041400,"lev":"Info","msg":"Only age.", "tags":{}, "fields":{"age":10}}
{"time":1651911648437041400,"lev":"Info","msg":"Only gender.", "tags":{}, "fields":{"gender":"male"}}
{"time":1651911648437565700,"lev":"Info","msg":"Ten.", "tags":{}, "fields":{"age":10}}
{"time":1651911676175910500,"lev":"Info","msg":"Twenty.","tags":{},"fields":{"age":20}}
{"time":1651911676175946300,"lev":"Error","msg":"Here's some error.","tags":{"session_id":"2147480001","user_id":"user_123"},"fields":{"the_error":"some-error"}}
```

## 特性

1. 使用链式操作，代码简洁。
2. 可对接其他日志库使用（实现对应的适配器`receiver`）。目前仅支持[seelog](https://github.com/cihub/seelog)。
3. 可添加两种额外信息，分别是`tags`和`fields`，均以键值对形式保存（同名字段会被覆盖）。其中`tags`必定是字符串，`fields`则支持多种类型。除此之外没有什么不同。
