package tail_dome
import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)
// tail demo
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen: true, //重新打开
		Follow: true,  //是否更随
		Location: &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件那个地方开始读
		MustExist: false,   //文件不存在不报错
		Poll: true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		msg *tail.Line
		ok bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}