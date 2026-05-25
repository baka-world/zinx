package znet

import "zinx/ziface"

type request struct {
	conn ziface.IConnection
	data []byte
}

func (r *request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *request) GetData() []byte {
	return r.data
}
