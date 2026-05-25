package ziface

// Wrap link and request info from client to Request
type IRequest interface {
	GetConn() IConnection
	GetData() []byte
}
