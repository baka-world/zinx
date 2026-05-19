package ziface

// 服务接口
type IServer interface {
	Start()
	Stop()
	Serve()
}
