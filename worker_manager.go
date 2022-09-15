package zworker

import (
	"sync"
)

var workerManagerOnce sync.Once
var workerManagerInstance *manager

func WorkerManager() *manager {
	workerManagerOnce.Do(func() {
		workerManagerInstance = newWorkerManager()
	})

	return workerManagerInstance
}

//单例
type manager struct {
	//当前WorkerManager监听的全部Worker工作组
	WorkerGroup sync.Map

	//workerID生成器 , TODO 目前用整型累加
	WorkerIdGen      int
	WorkerIdGenMutex *sync.Mutex
}

func newWorkerManager() *manager {
	instance := &manager{
		WorkerIdGen:      0,
		WorkerIdGenMutex: new(sync.Mutex),
	}

	return instance
}
