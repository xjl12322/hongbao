package tailog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"hongbao/logagent/kafka_go"
)
//var (
//	tailObj *tail.Tail
//	LogChan chan string
//)
//一个日志收集任务
type TailTask struct {
	path string
	topic string
	instance *tail.Tail
}

func NewTailTask(path,topic string)(tailObj *TailTask)  {
	tailObj = &TailTask{
		path:path,
		topic:topic,
	}
	tailObj.init() //更具路径打开对应的日志
	return
}


func (t *TailTask)init(){
	config := tail.Config{
		ReOpen: true, //重新打开
		Follow: true,  //是否更随
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件那个地方开始读
		MustExist: false,   //文件不存在不报错
		Poll: true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)

	}
	go t.run()  //直接采集日志发送到kafka

}


func (t *TailTask)run()  {
	for {
		select {
		case line := <- t.ReadChan():
			//3.2 发往kafka
			//kafka_go.SendToKafka(t.topic,line.Text)
			//3.2 发往kafka优化改为异步
			//先把日志发送到通道中
			kafka_go.SendToChan(t.topic,string(line))
			//kafka 那个包中有单独的groutime 去取日志数据发送到kafka


		}
	}
}


func (t *TailTask)ReadChan() <- chan *tail.Line{
	return t.instance.Lines
}


//func Init(fileName string)(err errors)  {
//	config := tail.Config{
//		ReOpen: true, //重新打开
//		Follow: true,  //是否更随
//		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件那个地方开始读
//		MustExist: false,   //文件不存在不报错
//		Poll: true,
//	}
//
//	tailObj, err = tail.TailFile(fileName, config)
//	if err != nil {
//		fmt.Println("tail file failed, err:", err)
//		return
//
//	}
//	return
//}



//func ReadChan() <- chan *tail.Line {
//	return tailObj.Lines
//
//}