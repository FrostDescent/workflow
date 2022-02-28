package workflow

import (
	"fmt"
	"time"
)

type WearUnderpantsAction struct {
}

func (a *WearUnderpantsAction) Run(i interface{}) {
	fmt.Println("我正在穿内裤...")
	time.Sleep(1 * time.Second)
}
