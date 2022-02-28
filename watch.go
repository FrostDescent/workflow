package workflow

import (
	"fmt"
	"time"
)

type WearWatchNodeAction struct {
}

func (a *WearWatchNodeAction) Run(i interface{}) {
	fmt.Println("我正在穿手表...")
	time.Sleep(1 * time.Second)
}
