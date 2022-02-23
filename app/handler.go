package app

type Handler interface {
	Handle() error
}
