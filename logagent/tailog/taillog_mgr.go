package tailog

import (
	"fmt"
	"hongbao/logagent/etcd"
	"time"
)
var tskMgr *tailLogMgr
type tailLogMgr struct {
	logEntryList []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry)  {
	tskMgr = &tailLogMgr{
		logEntryList:logEntryConf,//把当前的日志收集配置信息保存起来
		tskMap:make(map[string]*TailTask,16),
		newConfChan:make(chan []*etcd.LogEntry),  //无缓存区的通道
	}
	for _,logEntry := range logEntryConf{
		NewTailTask(logEntry.Path,logEntry.Topic)
	}
	go tskMgr.run()
}
//监听自己的newconfchan有了新的配置过来后旧做对应处理
// 1配置新增
// 2配置删除
// 3配置变更
func (t *tailLogMgr)run()  {
	for{
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新配置来乐",newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}
//向外暴漏一个函数，tskMgr 的newconfchan
func NewConfChan() <- chan []*etcd.LogEntry{
	return tskMgr.newConfChan
}

