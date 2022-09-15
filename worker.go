package zworker

import "fmt"

type Worker struct {
	//Worker Id
	Id int

	//Worker Name
	Name string

	//当前Worker正在处理的Message
	Msg Message

	//当前在处理的Message是否已经提交下游
	MsgCommitted bool

	//退出Worker的Channel
	Exit chan interface{}
}

//启动一个Worker业务
func (w *Worker) Start() (err error) {

	fmt.Printf("Worker ID %d , Name %s , Started...", w.Id, w.Name)

	//启动Worker
	for {
		select {
		//TODO 读取消息，并且调用业务逻辑
		}
	}

	return nil
}
