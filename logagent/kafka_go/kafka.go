package kafka_go

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)
type logData struct {
	topic string
	data string
}

//往kafka写i
//日志


//定义kafka 链接返回对象
var (
	client sarama.SyncProducer
	lodDataChan chan *logData
)
//INIT
func Init(addrs []string)(err error)  {
	config := sarama.NewConfig()
	//tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要
	//leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个
	//partition
	config.Producer.Return.Successes = true // 成功交付的消息将在
	//success channel返回
	// 构造⼀个消息
	//msg := &sarama.ProducerMessage{}
	//msg.Topic = "web_log"
	//msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	//初始化缓冲通道
	lodDataChan = make(chan *logData,1000000)
	//开启后台的goroutine 从发送数据到kafka里
	go sendToKafka()

	return

}
//给外部暴漏的函数，该函数只把日志数据发送到一个内部的channel
func SendToChan(topic,data string)  {
	msg := &logData{
		topic:topic,
		data:data,
	}
	lodDataChan <- msg

}
//往kafka发送数据
func sendToKafka() {
	for {
		select {
		case ld := <- lodDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 连接kafka
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			fmt.Println("xxxx")
			if err != nil {
				fmt.Println("send msg failed, err:", err)
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Microsecond*50)
		}


	}

}

























































































































































































































































































































































































































































































































































































































































































