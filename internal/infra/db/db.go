package db

type BaseDb interface {
	Disconnect() error
}
