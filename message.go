package zworker

import "context"

type Message struct {
	//数据源
	Src string
	//业务标识
	Busi string
	//数据种类
	Kind string
	//Message 唯一分布式ID
	ID string
	//详细数据
	Data interface{}
	//拓展数据
	Extra map[string]interface{}
	//上下文
	Ctx context.Context
}
