# logger

封装了底层实现的标准日志引擎

## 使用方法

### 初始化
```
config := logger.Config{
    EnableConsole: true,
    Level:         "debug",
    EnableFile:    true,
    FileLocation:  "/var/log/log.log",
    AppendCaller:  true,
}

err := logger.New(config, "logrus")
if err != nil {
    log.Fatalf("Could not instantiate log %s", err.Error())
}
```
#### 参数说明
|配置项|是否必须|可用值|说明|
|---|---|---|---|
|EnableConsole|否|true/false|是否将日志输出到控制台|
|Level|是|debug/info/warn/error/fatal/panic|日志级别|
|EnableFile|否|true/false|是否将日志输出到文件|
|FileLocation|否|`/var/log/log.log`|日志文件路径|
|AppendCaller|否|true/false|是否在日志中输出代码行数|

### 使用
#### 直接输出
```
logger.Debugf("...") // 输出debug级别的日志
logger.Infof("...") // 输出info级别的日志
logger.Warnf("...") // 输出warn级别的日志
logger.Errorf("...") // 输出error级别的日志
logger.Fatalf("...") // 输出fatal级别的日志，并终止
logger.Panicf("...") // 输出panic级别的日志，并终止
```

#### 带额外字段输出
```
contextLogger := logger.WithFields(logger.Fields{
    "carNo":      carNo,
    "requestURL": address,
})
contextLogger.Errorf("...")
```