package handler

type Breaker interface {
	Check() error
}
