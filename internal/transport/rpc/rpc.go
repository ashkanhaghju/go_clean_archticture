package rpc

type Rpc interface {
	Start(address string) error
	Shutdown() error
}
