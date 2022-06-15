# logcollection
日志实时采集
配置文件config.ini指定redis地址和tail采集的logpath,
logpath为文件的路径,可以指定多个,
实例:
多个logpath传入方式
logpath = D:/goproject/src/logagent/tail/tail.go,logpath = D:/goproject/src/logagent/tail/tail1.go,logpath = D:/goproject/src/logagent/tail/tail2.go
