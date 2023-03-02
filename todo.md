# todo

### todo
* 大文件分块上传
* auth模块crud
* Logger 按天分割 l.rotate()
  ```
  loggerWrite := &lumberjack.Logger{
    Filename: fileName, //文件名
    MaxSize: maxSize, //日志单文件的最大占用空间
    MaxAge: maxAge, //已经被分割存储的日志文件最大的留存时间，单位是天
    MaxBackups: maxBackup, //分割存储的日志文件最多的留存个数
    Compress: compress, //指定被分割之后的文件是否要压缩
    LocalTime: true,
    }

    //每日零点定时日志回滚分割实现时间上的分割
    if logType == "daily" {
        go  func() {
            for {
                nowTime := time.Now()
                nowTimeStr := nowTime.Format("2006-01-02")
                //使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
                t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
                // 第二天零点时间戳
                next := t2.AddDate(0, 0, 1)
                after := next.UnixNano() - nowTime.UnixNano() - 1
                <-time.After(time.Duration(after) * time.Nanosecond)
                loggerWrite.Rotate()
            }
        }()
    }
  ```