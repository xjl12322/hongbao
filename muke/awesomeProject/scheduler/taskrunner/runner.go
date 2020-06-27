package taskrunner

type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longLived bool  //是否从新创建还是用原来的
	Dispatcher fn
	Executor fn
}
//1 定义常量  defs
//2 构造函数
func NewRuner(size int,longlived bool,d fn,e fn)*Runner  {
	return &Runner{
		Controller:make(chan string,1),
		Error:make(chan string,1),
		Data:make(chan interface{},size),
		longLived:longlived,
		Dispatcher:d,
		Executor:e,
	}
}

func (r *Runner) startDispatch()  {
	defer func() {
		if !r.longLived{
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for{
		select {
		case c:= <-r.Controller:
			if c == READY_TO_DISPATCH{
				err:=r.Dispatcher(r.Data)
				if err !=nil{
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE{
				err := r.Executor(r.Data)
				if err != nil{
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e:= <-r.Error:
			if e == CLOSE{
				return
			}
		default:



		}
	}
}



func (r *Runner) StartAll()  {
	r.Controller <- READY_TO_DISPATCH  //防止组赛提前预支信号
	r.startDispatch()



}


















