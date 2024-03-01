package Demo

import (
	"context"
	"fmt"
)

type Demo struct {
	ctx context.Context
}

func NewDemo() *Demo {
	return &Demo{}
}

func (a *Demo) Start(ctx context.Context) {
	a.ctx = ctx
}

func (a *Demo) Atest(name string) string {
	return fmt.Sprintf("this is a test2!%s",name)
}