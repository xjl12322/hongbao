package main

import "sync"

type Raft struct {
	mu sync.Mutex //锁
	me int  //节点编号
	currentTerm int //当前任期
	votedFor int //那个节点投票
	state int  //3个状态  //0follower 1candidate 2 leader
	lastMessageTime int64 //发送最后一条数据时间
	currentLeader int  //设置i当前节点领导
	message chan bool //节点间发信息的通道
	electCh chan bool //选举通道
	heartBeat chan bool  //心跳信号的通道
	heartbeatRe chan bool //返回心跳信号的通道
	timeout int  //  超市时间
}
//定义3节点的常量
const raftCount = 3
//声明leader对象
type Leader struct {
	//任期
	Term int
	//编号
	LeaderId int
}
//0 还没上任  -1 没有编号
var leader = Leader{}
func main()  {
	//有3哥节点 都是follower
	//若有canidate状态 进行投票拉票
	//产生leader

	//创建3个节点
	for i:= 0;i<raftCount;i++{
		//
	}

	for {;

	}



}
//创建节点对象
func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	rf.votedFor = -1
	//0 follower ,1 candidate ,2 leader
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	rf.setTerm(0)

}

