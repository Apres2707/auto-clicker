package service

import (
	"fmt"
	"github.com/mattn/go-tty"
)

const termSignal = "q"

type Breaker struct{}

func (c Breaker) Check() error {
	ttyItem, err := tty.Open()
	if err != nil {
		return fmt.Errorf("tty open: %w", err)
	}
	defer func() { _ = ttyItem.Close() }()

	for {
		r, err := ttyItem.ReadRune()
		if err != nil {
			return fmt.Errorf("read rune: %w", err)
		}
		if string(r) == termSignal {
			return nil
		}
	}
}
