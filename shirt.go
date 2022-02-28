package workflow

import (
	"fmt"
	"time"
)

type WearShirtNodeAction struct {
}

func (a *WearShirtNodeAction) Run(i interface{}) {
	fmt.Println("我正在穿T恤...")
	time.Sleep(1 * time.Second)
}
