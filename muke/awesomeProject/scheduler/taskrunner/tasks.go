package taskrunner


////延时删除功能
import (
	"errors"
	//
	"hongbao/muke/awesomeProject/scheduler/dbops"
	"os"
	"sync"
	//
	"log"
)
//从数据库里读要删除的信息
func deleteVideo(vid string)error  {
	err := os.Remove(VIDEO_PATH+vid)
	if err != nil && !os.IsNotExist(err){
		log.Printf("deleting video error : %v",err)
		return err
	}
	return nil
}
//
//
func VideoClearDispatcher(dc dataChan)error  {
	//从要删除表里读取要删除的视频id
	res,err := dbops.ReadVideoDeletionRecord(3)//每次读多少
	if err != nil{
		log.Printf("video clear dispathcher err:%v",err)
		return err
	}
	if len(res) == 0{  //表示没有任何数据读取出来
		return errors.New("all tasks finished")
	}

	for _,id := range res{
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan)error  {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
			case vid :=<- dc:
				go func(id interface{}) {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}








