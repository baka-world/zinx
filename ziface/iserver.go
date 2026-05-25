package ziface

// IServer 服务接口
type IServer interface {
	Start()
	Stop()
	Serve()
	// AddRouter Register router
	AddRouter(router IRouter)
}
